/*
 * @Author: NyanCatda
 * @Date: 2022-11-23 14:46:04
 * @LastEditTime: 2022-11-23 19:21:34
 * @LastEditors: NyanCatda
 * @Description: 连接到服务端
 * @FilePath: \Atsuko\Connect.go
 */
package main

import (
	"fmt"
	"net"

	"github.com/nyancatda/Atsuko/internal/MessageProcessing"
	"github.com/nyancatda/Atsuko/internal/TCPComm"
)

/**
 * @description: 连接到服务端
 * @param {string} Address 服务端地址
 * @return {error} 错误信息
 */
func Connect(Address string) {
	go TCPComm.StartClient(Address, MessageChan, func(Msg string, Conn net.Conn) {
		// 解析消息
		Content, err := MessageProcessing.Receive(Msg)
		if err != nil {
			fmt.Println("消息解析失败:", err)
			return
		}

		fmt.Println(Conn.RemoteAddr().String() + "> " + Content)
	})
}
