package cryptControllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func EncryptAES(plainData, secret []byte) ([]byte, error) {
	block, _ := aes.NewCipher(secret)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, err
	}
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return []byte{}, err
	}
	cipherData := gcm.Seal(nonce, nonce, plainData, nil)
	return cipherData, nil
}

func DecryptAES(cipherData, secret []byte) ([]byte, error) {
	block, err := aes.NewCipher(secret)
	if err != nil {
		return []byte{}, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, err
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := cipherData[:nonceSize], cipherData[nonceSize:]
	plainData, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return []byte{}, err
	}
	return plainData, nil
}

func WriteFile(filename string, filedata []byte) error {
	f, err := os.Create("./temp/" + filename)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = f.Write(filedata)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
