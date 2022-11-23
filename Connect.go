/*
 * @Author: NyanCatda
 * @Date: 2022-11-23 14:46:04
 * @LastEditTime: 2022-11-23 15:24:00
 * @LastEditors: NyanCatda
 * @Description: 连接到服务端
 * @FilePath: \Atsuko\Connect.go
 */
package main

import (
	"fmt"

	"github.com/nyancatda/Atsuko/internal/Command"
	"github.com/nyancatda/Atsuko/internal/TCPComm"
)

func Connect(CommandStr string) {
	_, Parameter := Command.Parse(CommandStr)

	if len(Parameter) != 1 {
		fmt.Println("命令格式错误，应为：connect [Host:Port]")
		return
	}

	go TCPComm.StartClient(Parameter[0], MessageChan, func(Msg string) {
		// 接收消息回调
		fmt.Println("\n" + Msg)

		// 打印命令提示符
		fmt.Print("\r>")
	})
}
