/*
 * @Author: NyanCatda
 * @Date: 2022-11-23 14:49:20
 * @LastEditTime: 2022-11-23 19:51:38
 * @LastEditors: NyanCatda
 * @Description: 启动服务端
 * @FilePath: \Atsuko\StartServe.go
 */
package main

import (
	"context"
	"fmt"
	"net"

	"github.com/nyancatda/Atsuko/internal/Flag"
	"github.com/nyancatda/Atsuko/internal/MessageProcessing"
	"github.com/nyancatda/Atsuko/internal/TCPComm"
)

func StartServe(Context context.Context) {
	go TCPComm.StartServe(Flag.Flag.ListenPort, Context, MessageChan, func(Msg string, Conn net.Conn) {
		// 解析消息
		Content, err := MessageProcessing.Receive(Msg)
		if err != nil {
			fmt.Println("消息解析失败:", err)
			return
		}

		fmt.Println(Conn.RemoteAddr().String() + "> " + Content)
	})
}
