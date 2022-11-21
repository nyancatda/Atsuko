/*
 * @Author: NyanCatda
 * @Date: 2022-11-21 21:16:21
 * @LastEditTime: 2022-11-21 21:38:05
 * @LastEditors: NyanCatda
 * @Description: 公钥部分
 * @FilePath: \Atsuko\internal\EncryptDecrypt\PublicKey.go
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
 * @description: 使用公钥加密数据
 * @param {SecretKey.Key} Key 密钥对
 * @param {[]byte} Body 明文
 * @return {string} 密文
 * @return {error} 错误
 */
func Encrypt(Key SecretKey.Key, Body []byte) (string, error) {
	EncryptData, err := rsa.EncryptPKCS1v15(rand.Reader, Key.PublicKey, Body)
	if err != nil {
		return "", err
	}

	// base64编码
	EncryptBody := base64.StdEncoding.EncodeToString(EncryptData)

	return EncryptBody, nil
}

/**
 * @description: 验证签名
 * @param {SecretKey.Key} Key 密钥对
 * @param {[]byte} Body 原始消息
 * @param {string} Sign 签名
 * @return {bool} 是否正确
 */
func VerifySign(Key SecretKey.Key, Body []byte, Sign string) bool {
	//计算散列值
	Hash := sha256.New()
	Hash.Write(Body)
	BodyHashBytes := Hash.Sum(nil)

	// base64解码签名
	SignBody, err := base64.StdEncoding.DecodeString(Sign)
	if err != nil {
		return false
	}

	// 验证签名
	err = rsa.VerifyPKCS1v15(Key.PublicKey, crypto.SHA256, BodyHashBytes, SignBody)

	return err == nil
}
