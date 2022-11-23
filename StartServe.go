/*
 * @Author: NyanCatda
 * @Date: 2022-11-23 14:49:20
 * @LastEditTime: 2022-11-23 15:55:56
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
	"github.com/nyancatda/Atsuko/internal/TCPComm"
)

func StartServe(Context context.Context) {
	go TCPComm.StartServe(Flag.Flag.ListenPort, Context, MessageChan, func(Msg string, Conn net.Conn) {
		// 接收消息回调
		fmt.Println(Conn.RemoteAddr().String() + "> " + Msg)
	})
}
