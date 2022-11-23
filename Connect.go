/*
 * @Author: NyanCatda
 * @Date: 2022-11-23 14:46:04
 * @LastEditTime: 2022-11-23 15:38:56
 * @LastEditors: NyanCatda
 * @Description: 连接到服务端
 * @FilePath: \Atsuko\Connect.go
 */
package main

import (
	"fmt"

	"github.com/nyancatda/Atsuko/internal/TCPComm"
)

/**
 * @description: 连接到服务端
 * @param {string} Address 服务端地址
 * @return {error} 错误信息
 */
func Connect(Address string) {
	go TCPComm.StartClient(Address, MessageChan, func(Msg string) {
		// 接收消息回调
		fmt.Println("\n" + Msg)

		// 打印命令提示符
		fmt.Print("\r>")
	})
}
