/*
 * @Author: NyanCatda
 * @Date: 2022-11-23 17:50:48
 * @LastEditTime: 2022-11-23 19:50:37
 * @LastEditors: NyanCatda
 * @Description: 消息处理
 * @FilePath: \Atsuko\internal\MessageProcessing\MessageProcessing.go
 */
package MessageProcessing

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/nyancatda/Atsuko/internal/EncryptDecrypt"
	"github.com/nyancatda/Atsuko/internal/SecretKey"
)

var (
	MyKey     SecretKey.Key
	Personkey SecretKey.Key
)

type Content struct {
	Message string
	Sign    string
}

func init() {
	// 获取我的密钥
	if MyKey == (SecretKey.Key{}) {
		Key, err := SecretKey.ReadKey("./")
		if err != nil {
			fmt.Println(err)
			return
		}
		MyKey = Key
	}
}

/**
 * @description: 发送消息处理
 * @param {string} Message 消息
 * @return {string} 消息内容
 * @return {error} 错误信息
 */
func Send(Message string) (string, error) {
	// 加密消息
	EncryptMessage, err := EncryptDecrypt.Encrypt(Personkey, []byte(Message))
	if err != nil {
		return "", err
	}

	// 生成签名
	Sign, err := EncryptDecrypt.Sign(MyKey, []byte(Message))
	if err != nil {
		return "", err
	}

	// 生成消息
	Msg := Content{
		Message: EncryptMessage,
		Sign:    Sign,
	}

	// 转换为json
	Content, err := json.Marshal(Msg)
	if err != nil {
		return "", err
	}

	return string(Content), nil
}

/**
 * @description: 接收消息处理
 * @param {string} MessageContent 消息Json
 * @return {string} 消息内容
 * @return {error} 错误信息
 */
func Receive(MessageContent string) (string, error) {
	// 解析json
	var Msg Content
	err := json.Unmarshal([]byte(MessageContent), &Msg)
	if err != nil {
		return "", err
	}

	// 解析消息
	Message, err := EncryptDecrypt.Decrypt(MyKey, Msg.Message)
	if err != nil {
		return "", err
	}

	// 验证签名
	SignOK := EncryptDecrypt.VerifySign(Personkey, []byte(Message), Msg.Sign)
	if !SignOK {
		return "", errors.New("签名验证失败")
	}

	return string(Message), nil
}
