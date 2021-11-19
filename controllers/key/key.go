package keyControllers

import (
	"Github/sarthakpranesh/silvershare/connections"
	"errors"

	"gorm.io/gorm"
)

type Key struct {
	gorm.Model
	Secret     []byte   `json:"secret"`
	UserId     string   `json:"uid"`
	SharedWith []Shared `json:"sharedWith"`
}

type Shared struct {
	KeyId  uint   `json:"keyId"`
	UserId string `json:"uid"`
	Email  string `json:"email"`
}

func CreateKey(k *Key) error {
	db := connections.DB
	result := db.Create(k)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AllUserKeys(uid string) (*[]Key, error) {
	db := connections.DB
	var keys []Key
	result := db.Table("keys").Where("user_id", uid).Find(&keys)
	if result.Error != nil {
		return &keys, result.Error
	}
	return &keys, nil
}

func GetKey(id uint, uid string) (*Key, error) {
	db := connections.DB
	var key Key
	result := db.Table("keys").Where(id).First(&key)
	if result.Error != nil {
		return &key, result.Error
	}
	var shareds []Shared
	result = db.Table("shareds").Where(id).Find(&shareds)
	if result.Error != nil {
		return &key, result.Error
	}
	key.SharedWith = shareds
	if key.UserId != uid {
		var isSharedKey bool
		for _, s := range key.SharedWith {
			if s.UserId == uid {
				isSharedKey = true
				break
			}
		}
		if !isSharedKey {
			return &Key{}, errors.New("not authorized")
		}
	}
	return &key, nil
}

func ShareKey(uid string, share Shared) error {
	db := connections.DB
	var key Key
	result := db.Table("keys").Where(share.KeyId).First(&key)
	if result.Error != nil {
		return result.Error
	}
	if key.UserId != uid {
		return errors.New("not authorized")
	}
	var userEmail struct {
		Uid string `json:"uid"`
	}
	result = db.Table("users").Where("email", share.Email).First(&userEmail)
	if result.Error != nil {
		return result.Error
	}
	if userEmail.Uid == "" {
		return errors.New("no user found")
	}
	share.UserId = userEmail.Uid
	result = db.Create(share)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
