/*
 * @Author: NyanCatda
 * @Date: 2022-11-23 15:44:28
 * @LastEditTime: 2022-11-23 15:46:57
 * @LastEditors: NyanCatda
 * @Description:
 * @FilePath: \Atsuko\Send.go
 */
package main

import (
	"fmt"

	"github.com/nyancatda/Atsuko/internal/Command"
	"github.com/nyancatda/Atsuko/internal/TCPComm"
)

/**
 * @description: 发送消息处理
 * @param {string} CommandStr 命令字符串
 * @return {*}
 */
func SendMessage(CommandStr string) {
	// 判断连接状态标识
	if !TCPComm.ConnectionStatus {
		fmt.Println("请先建立连接")
		return
	}

	// 解析命令参数
	_, Parameter := Command.Parse(CommandStr)
	if len(Parameter) != 1 {
		fmt.Println("命令格式错误，应为：send [Message...]")
		return
	}

	// 拼接所有参数组成消息
	var Message string
	for _, Value := range Parameter {
		Message += Value
	}

	// 发送消息
	MessageChan <- Message
}
