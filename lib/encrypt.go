package lib

import (
	"bytes"
	"crypto/des"
	"encoding/base64"
	"errors"

	"github.com/juxuny/clipboard-sync/lib/env"
)

var (
	DESKey = env.GetString(env.Key.EncryptSecretKey, "")
)

/**
 * DES加密算法
 */
func DESEncrypt(dst string) string {
	out, _ := desECBEncrypt([]byte(dst), []byte(DESKey))
	return base64.StdEncoding.EncodeToString(out)
}

/**
 * DES解密算法
 */
func DESDecrypt(dst string) string {
	data, _ := base64.StdEncoding.DecodeString(dst)
	out, _ := desECBDecrypt(data, []byte(DESKey))
	return string(out)
}

/**
 * DES解密算,自定义key
 */
func DESDecryptByKey(dst string, desKey string) string {
	data, _ := base64.StdEncoding.DecodeString(dst)
	out, _ := desECBDecrypt(data, []byte(desKey))
	return string(out)
}

func desECBEncrypt(data, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	data = mPKCS5Padding(data, bs)
	if len(data)%bs != 0 {
		return nil, errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return out, nil
}

func desECBDecrypt(data, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = mPKCS5UnPadding(out)
	return out, nil
}

func mPKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func mPKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
