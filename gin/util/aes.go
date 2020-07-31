package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func Encrypt(cryptkey string,original_word string) (string){
	var aeskey = []byte(cryptkey)
	var original_data = []byte(original_word)
	encrypted_data,err := AesEncrypt(original_data,aeskey)
	if err != nil {
		fmt.Println(err)
	}

	ciphertext := base64.StdEncoding.EncodeToString(encrypted_data)
	return ciphertext
}

func Decrypt(cryptkey string,ciphertext string)(string){
	var aeskey = []byte(cryptkey)
	cipher_data,err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		fmt.Println(err)
	}
	data,err:=AesDecrypt(cipher_data,aeskey)
	if err != nil {
		fmt.Println(err)
	}
	plaintext := base64.StdEncoding.EncodeToString(data)
	return plaintext
}
