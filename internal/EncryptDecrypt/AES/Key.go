/*
 * @Author: NyanCatda
 * @Date: 2022-11-24 11:57:22
 * @LastEditTime: 2022-11-24 12:00:38
 * @LastEditors: NyanCatda
 * @Description: 生成随机密钥
 * @FilePath: \Atsuko\internal\EncryptDecrypt\AES\Key.go
 */
package AES

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

/**
 * @description: 生成随机密钥
 * @return {string} 密钥
 */
func CreateKey() (string, error) {
	// 随机生成32字节的密钥
	Key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, Key); err != nil {
		return "", err
	}

	// 使用base64编码密钥
	return base64.StdEncoding.EncodeToString(Key), nil
}

/**
 * @description: 解析base64密钥
 * @param {string} Key 密钥
 * @return {[]byte} 解析后的密钥
 * @return {error} 错误
 */
func parseKey(Key string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(Key)
}
