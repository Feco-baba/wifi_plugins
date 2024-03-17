package utils

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
)

func FileContentDecrypt(fileContent string) (string, error) {
	key := []byte("LpnEvVc73J7mz8i2tuzRC4g0ByvpYEEX") // AES-256要求32字节密钥
	serversSubscribedURLByteArr, err := hex.DecodeString(string(fileContent))
	if err != nil {
		return "", err
	}
	decryptedStr, err := Decrypt(serversSubscribedURLByteArr, key)
	if err != nil {
		return "", err
	}
	return string(decryptedStr), nil
}

// decrypt解密数据
func Decrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(data)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("ciphertext is not a multiple of the block size")
	}
	decrypted := make([]byte, len(data))
	for bs, be := 0, block.BlockSize(); bs < len(data); bs, be = bs+block.BlockSize(), be+block.BlockSize() {
		block.Decrypt(decrypted[bs:be], data[bs:be])
	}
	return Unpad(decrypted), nil
}

// unpad 删除PKCS#7填充
func Unpad(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
