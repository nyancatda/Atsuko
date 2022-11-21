/*
 * @Author: NyanCatda
 * @Date: 2022-11-21 20:13:49
 * @LastEditTime: 2022-11-21 20:55:40
 * @LastEditors: NyanCatda
 * @Description: 创建密钥对
 * @FilePath: \Atsuko\internal\SecretKey\Create.go
 */
package SecretKey

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

type Key struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

/**
 * @description: 创建密钥对
 * @return {Key} 密钥对
 * @return {error} 错误
 */
func CreateKey() (Key, error) {
	// 生成私钥
	PrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return Key{}, err
	}
	// 生成公钥
	PublicKey := &PrivateKey.PublicKey

	// 保存密钥对
	Key := Key{
		PrivateKey: PrivateKey,
		PublicKey:  PublicKey,
	}
	return Key, nil
}

/**
 * @description: PEM编码私钥
 * @return {[]byte} 私钥
 */
func (Key Key) PEMEncodePrivateKey() []byte {
	// 将RSA私钥序列化为ASN.1 PKCS#1 DER编码
	PrivateKeyData := x509.MarshalPKCS1PrivateKey(Key.PrivateKey)

	// 将私匙做PEM数据编码
	PrivateKey := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: PrivateKeyData,
	})

	return PrivateKey
}

/**
 * @description: PEM编码公钥
 * @return {[]byte} 公钥
 * @return {error} 错误
 */
func (Key Key) PEMEncodePublicKey() ([]byte, error) {
	// 将RSA公钥序列化为ASN.1 PKCS#1 DER编码
	PublicKeyData := x509.MarshalPKCS1PublicKey(Key.PublicKey)

	// 将公匙做PEM数据编码
	PublicKey := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: PublicKeyData,
	})

	return PublicKey, nil
}
