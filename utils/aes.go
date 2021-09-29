package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// AESCBCEncrypt AES CBC 加密
/*
   AES  CBC 加密
   key:加密key
   plaintext：加密明文
   ciphertext:解密返回字节字符串[ 整型以十六进制方式显示]

*/
func AESCBCEncrypt(key, plaintext string) (ciphertext string) {
	plainbyte := []byte(plaintext)
	keybyte := []byte(key)
	if len(plainbyte)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}
	block, err := aes.NewCipher(keybyte)
	if err != nil {
		panic(err)
	}

	cipherbyte := make([]byte, aes.BlockSize+len(plainbyte))
	iv := cipherbyte[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherbyte[aes.BlockSize:], plainbyte)

	ciphertext = fmt.Sprintf("%x\n", cipherbyte)
	return
}

// AESCBCDecrypter ASE CBC 解密
/*
   AES  CBC 解码
   key:解密key
   ciphertext:加密返回的串
   plaintext：解密后的字符串
*/
func AESCBCDecrypter(key, ciphertext string) (plaintext string) {
	cipherbyte, _ := hex.DecodeString(ciphertext)
	keybyte := []byte(key)
	block, err := aes.NewCipher(keybyte)
	if err != nil {
		panic(err)
	}
	if len(cipherbyte) < aes.BlockSize {
		panic("ciphertext too short")
	}

	iv := cipherbyte[:aes.BlockSize]
	cipherbyte = cipherbyte[aes.BlockSize:]
	if len(cipherbyte)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherbyte, cipherbyte)

	//fmt.Printf("%s\n", ciphertext)
	plaintext = string(cipherbyte[:])
	return
}

// AESGCMEncrypt ASE GCM 加密
/*
   AES  GCM 加密
   key:加密key
   plaintext：加密明文
   ciphertext:解密返回字节字符串[ 整型以十六进制方式显示]

*/
func AESGCMEncrypt(key, plaintext string) (ciphertext, noncetext string) {
	plainbyte := []byte(plaintext)
	keybyte := []byte(key)
	block, err := aes.NewCipher(keybyte)
	if err != nil {
		panic(err.Error())
	}

	// 由于存在重复的风险，请勿使用给定密钥使用超过2^32个随机值。
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	cipherbyte := aesgcm.Seal(nil, nonce, plainbyte, nil)
	ciphertext = fmt.Sprintf("%x\n", cipherbyte)
	noncetext = fmt.Sprintf("%x\n", nonce)
	return
}

// AESGCMDecrypter AES GCM 解密
/*
   AES  GCM 解码
   key:解密key
   ciphertext:加密返回的串
   plaintext：解密后的字符串
*/
func AESGCMDecrypter(key, ciphertext, noncetext string) (plaintext string) {
	cipherbyte, _ := hex.DecodeString(ciphertext)
	nonce, _ := hex.DecodeString(noncetext)
	keybyte := []byte(key)
	block, err := aes.NewCipher(keybyte)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plainbyte, err := aesgcm.Open(nil, nonce, cipherbyte, nil)
	if err != nil {
		panic(err.Error())
	}

	//fmt.Printf("%s\n", ciphertext)
	plaintext = string(plainbyte[:])
	return
}
