/*
 * @Author: NyanCatda
 * @Date: 2022-11-21 20:24:55
 * @LastEditTime: 2022-11-23 16:46:49
 * @LastEditors: NyanCatda
 * @Description: 解析PEM密钥对
 * @FilePath: \Atsuko\internal\SecretKey\Parsing.go
 */
package SecretKey

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
)

/**
 * @description: 解析PEM编码的公钥
 * @param {[]byte} PublicKey PEM编码的公钥
 * @return {*rsa.PublicKey} 公钥
 */
func PEMDecodePublicKey(PublicKey []byte) (Key, error) {
	// PEM解码
	PublicKeyBlock, _ := pem.Decode(PublicKey)
	if PublicKeyBlock == nil {
		return Key{}, errors.New("public key error")
	}
	// DER解码，获得公钥对象
	PublicKeyStruct, err := x509.ParsePKCS1PublicKey(PublicKeyBlock.Bytes)
	if err != nil {
		return Key{}, err
	}

	Key := Key{
		PublicKey: PublicKeyStruct,
	}

	return Key, nil
}

/**
 * @description: 解析PEM编码的私钥
 * @param {[]byte} PrivateKey PEM编码的私钥
 * @return {error} 错误
 */
func PEMDecodePrivateKey(PrivateKey []byte) (Key, error) {
	// PEM解码
	PrivateKeyBlock, _ := pem.Decode(PrivateKey)
	if PrivateKeyBlock == nil {
		return Key{}, errors.New("private key error")
	}
	// DER解码，获得私钥对象
	PrivateKeyStruct, err := x509.ParsePKCS1PrivateKey(PrivateKeyBlock.Bytes)
	if err != nil {
		return Key{}, err
	}

	Key := Key{
		PrivateKey: PrivateKeyStruct,
	}

	return Key, nil
}
