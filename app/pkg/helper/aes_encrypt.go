package helper

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
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
	iv := key[:blockSize]
	cipherText := make([]byte, len(payload))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText, payload)

	return cipherText, nil
}

func aesCBCDecrypt(encryptStr []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := key[:block.BlockSize()]
	mode := cipher.NewCBCDecrypter(block, iv)
	cipherText := make([]byte, len(encryptStr))
	mode.CryptBlocks(cipherText, encryptStr)
	encryptStr = pKCS7UnPadding(cipherText)

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
