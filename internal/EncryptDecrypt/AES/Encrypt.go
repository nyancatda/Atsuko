/*
 * @Author: NyanCatda
 * @Date: 2022-11-24 11:28:17
 * @LastEditTime: 2022-11-24 12:01:02
 * @LastEditors: NyanCatda
 * @Description: AES加密封装
 * @FilePath: \Atsuko\internal\EncryptDecrypt\AES\Encrypt.go
 */
package AES

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

/**
 * @description: AES加密
 * @param {[]byte} Key 密钥
 * @param {[]byte} Data 明文
 * @return {string} 密文 (base64)
 * @return {error} 错误
 */
func Encrypt(Key string, Data []byte) (string, error) {
	// 解析密钥
	KeyByte, err := parseKey(Key)

	Block, err := aes.NewCipher(KeyByte)
	if err != nil {
		return "", err
	}
	// PKCS7填充加密块
	BlockSize := Block.BlockSize()
	EncryptBytes := PKCS7Padding(Data, BlockSize)

	Encrypt := make([]byte, len(EncryptBytes))
	// 使用CBC加密模式
	BlockMode := cipher.NewCBCEncrypter(Block, KeyByte[:BlockSize])
	BlockMode.CryptBlocks(Encrypt, EncryptBytes)

	// 使用base64编码
	Base64EncryptData := base64.StdEncoding.EncodeToString(Encrypt)

	return Base64EncryptData, nil
}

/**
 * @description: PKCS7填充
 * @param {[]byte} Data 需要填充的数据
 * @param {int} BlockSize 块大小
 * @return {[]byte} 填充后的数据
 */
func PKCS7Padding(Data []byte, BlockSize int) []byte {
	//判断缺少几位长度
	Padding := BlockSize - len(Data)%BlockSize
	//补足位数
	PadText := bytes.Repeat([]byte{byte(Padding)}, Padding)
	return append(Data, PadText...)
}
