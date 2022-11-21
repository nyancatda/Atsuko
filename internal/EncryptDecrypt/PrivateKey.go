/*
 * @Author: NyanCatda
 * @Date: 2022-11-21 20:56:49
 * @LastEditTime: 2022-11-21 21:36:40
 * @LastEditors: NyanCatda
 * @Description: 私钥部分
 * @FilePath: \Atsuko\internal\EncryptDecrypt\PrivateKey.go
 */
package EncryptDecrypt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"

	"github.com/nyancatda/Atsuko/internal/SecretKey"
)

/**
 * @description: 使用私钥解密数据
 * @param {SecretKey.Key} Key 密钥对
 * @param {string} EncryptBody 密文
 * @return {[]byte} 明文
 */
func Decrypt(Key SecretKey.Key, EncryptBody string) ([]byte, error) {
	// base64解码密文
	EncryptData, err := base64.StdEncoding.DecodeString(EncryptBody)

	DecryptData, err := rsa.DecryptPKCS1v15(rand.Reader, Key.PrivateKey, EncryptData)
	if err != nil {
		return nil, err
	}

	return DecryptData, nil
}

/**
 * @description: 使用私钥签名消息
 * @param {SecretKey.Key} Key 密钥对
 * @param {[]byte} Body 消息
 * @return {string} 签名
 * @return {error} 错误
 */
func Sign(Key SecretKey.Key, Body []byte) (string, error) {
	//计算散列值
	Hash := sha256.New()
	Hash.Write(Body)
	BodyHashBytes := Hash.Sum(nil)

	// 签名
	Signature, err := rsa.SignPKCS1v15(rand.Reader, Key.PrivateKey, crypto.SHA256, BodyHashBytes)
	if err != nil {
		return "", err
	}

	// base64编码签名
	SignBody := base64.StdEncoding.EncodeToString(Signature)

	return SignBody, nil
}
