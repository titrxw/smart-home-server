package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

func pKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func aesCBCEncrypt(payload []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	payload = pKCS7Padding(payload, blockSize)
	cipherText := make([]byte, blockSize+len(payload))
	iv := cipherText[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[blockSize:], payload)

	return cipherText, nil
}

func aesCBCDecrypt(encryptStr []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if len(encryptStr) < blockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := encryptStr[:blockSize]
	encryptStr = encryptStr[blockSize:]
	if len(encryptStr)%blockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encryptStr, encryptStr)
	encryptStr = pKCS7UnPadding(encryptStr)

	return encryptStr, nil
}

func Encrypt(payload string, key string) (string, error) {
	data, err := aesCBCEncrypt([]byte(payload), []byte(key))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func Decrypt(encryptStr string, key string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encryptStr)
	if err != nil {
		return "", err
	}

	payload, err := aesCBCDecrypt(data, []byte(key))
	if err != nil {
		return "", err
	}

	return string(payload), nil
}
