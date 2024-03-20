package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func FileContentDecrypt(fileContent string) (string, error) {
	key := []byte("LpnEvVc73J7mz8i2tuzRC4g0ByvpYEEX") // AES-256要求32字节密钥
	decryptedStr, err := Decrypt(fileContent, key)
	if err != nil {
		return "FileContentDecrypt", err
	}
	return string(decryptedStr), nil
}

// PKCS7Padding 填充函数
func PKCS7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

// PKCS7UnPadding 去填充函数
func PKCS7UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

// Encrypt 使用AES-256 CBC模式加密字符串
func Encrypt(message, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// 对明文进行PKCS7填充
	padded := PKCS7Padding(message, block.BlockSize())

	mode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(padded))
	mode.CryptBlocks(ciphertext, padded)

	// 将IV和密文一起返回，以便解密时使用
	return base64.StdEncoding.EncodeToString(append(iv, ciphertext...)), nil
}

// Decrypt 使用AES-256 CBC模式解密字符串
func Decrypt(encrypted string, key []byte) ([]byte, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	// 去除PKCS7填充
	unpadded := PKCS7UnPadding(plaintext)

	return unpadded, nil
}
