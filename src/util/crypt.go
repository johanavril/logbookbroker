package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type cipherKey struct {
	Value string `yaml:"cipher_key"`
}

func getKey() (*cipherKey, error) {
	yamlFile, err := ioutil.ReadFile("../config/cipher.yml")
	if err != nil {
		return nil, err
	}

	c := cipherKey{}

	if err := yaml.Unmarshal(yamlFile, &c); err != nil {
		return nil, err
	}

	return &c, nil
}

func Encrypt(unencrypted string) (string, error) {
	plaintext := []byte(unencrypted)
	key, err := getKey()
	if err != nil {
		return "", err
	}

	byteKey := []byte(key.Value)
	c, err := aes.NewCipher(byteKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	return fmt.Sprintf("%x", ciphertext), nil
}

func Decrypt(encrypted string) (string, error) {
	ciphertext, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", nil
	}

	key, err := getKey()
	if err != nil {
		return "", err
	}

	byteKey := []byte(key.Value)
	c, err := aes.NewCipher(byteKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if nonceSize > len(ciphertext) {
		return "", errors.New("Invalid Nonce.")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
