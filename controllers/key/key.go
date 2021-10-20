package keyControllers

import (
	"Github/sarthakpranesh/silvershare/connections"
	"errors"

	"gorm.io/gorm"
)

type Key struct {
	gorm.Model
	Secret     []byte   `json:"secret"`
	UserId     uint     `json:"userId"`
	SharedWith []Shared `json:"sharedWith"`
}

type Shared struct {
	KeyId  uint `json:"keyId"`
	UserId uint `json:"userId"`
}

func CreateKey(k *Key) error {
	db, _ := connections.PostgresConnector()
	db.AutoMigrate(&Key{})
	result := db.Create(k)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AllUserKeys(UserId uint) (*[]Key, error) {
	db, _ := connections.PostgresConnector()
	db.AutoMigrate(&Key{})
	var keys []Key
	result := db.Table("keys").Where("user_id", UserId).Find(&keys)
	if result.Error != nil {
		return &keys, result.Error
	}
	return &keys, nil
}

func GetKey(id uint, user_id uint) (*Key, error) {
	db, _ := connections.PostgresConnector()
	db.AutoMigrate(&Key{})
	var key Key
	result := db.Table("keys").Where(id).First(&key)
	if result.Error != nil {
		return &key, result.Error
	}
	if key.UserId != user_id {
		var shareds []Shared
		result := db.Table("shareds").Where(id).Find(&shareds)
		if result.Error != nil {
			return &key, result.Error
		}
		var isSharedKey bool
		for _, s := range shareds {
			if s.UserId == user_id {
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

func ShareKey(user_id uint, share Shared) error {
	db, _ := connections.PostgresConnector()
	db.AutoMigrate(&Shared{})
	var key Key
	result := db.Table("keys").Where(share.KeyId).First(&key)
	if result.Error != nil {
		return result.Error
	}
	if key.UserId != user_id {
		return errors.New("not authorized")
	}
	result = db.Create(share)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
