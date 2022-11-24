/*
 * @Author: NyanCatda
 * @Date: 2022-11-24 11:37:23
 * @LastEditTime: 2022-11-24 12:01:51
 * @LastEditors: NyanCatda
 * @Description: AES解密封装
 * @FilePath: \Atsuko\internal\EncryptDecrypt\AES\Decrypt.go
 */
package AES

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

/**
 * @description: AES解密
 * @param {[]byte} Key 密钥
 * @param {[]byte} Data 密文
 * @return {[]byte} 明文
 * @return {error} 错误
 */
func Decrypt(Key string, Data string) ([]byte, error) {
	// 解析密钥
	KeyByte, err := parseKey(Key)

	// 使用base64解码密文
	EncryptData, err := base64.StdEncoding.DecodeString(Data)
	if err != nil {
		return nil, err
	}

	Block, err := aes.NewCipher(KeyByte)
	if err != nil {
		return nil, err
	}

	BlockSize := Block.BlockSize()
	DecryptData := make([]byte, len(EncryptData))
	// 使用CBC模式解密
	BlockMode := cipher.NewCBCDecrypter(Block, KeyByte[:BlockSize])
	BlockMode.CryptBlocks(DecryptData, EncryptData)

	// 去除填充
	DecryptData, err = PKCS7UNPadding(DecryptData)
	if err != nil {
		return nil, err
	}
	return DecryptData, nil
}

/**
 * @description: 去除PKCS7填充
 * @param {[]byte} Data 需要去除填充的数据
 * @return {[]byte} 去除填充后的数据
 * @return {error} 错误
 */
func PKCS7UNPadding(Data []byte) ([]byte, error) {
	Length := len(Data)
	if Length == 0 {
		return nil, errors.New("Data is empty")
	}
	//获取填充的个数
	UNPadding := int(Data[Length-1])
	return Data[:(Length - UNPadding)], nil
}
