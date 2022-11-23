/*
 * @Author: NyanCatda
 * @Date: 2022-11-23 14:49:20
 * @LastEditTime: 2022-11-23 15:12:57
 * @LastEditors: NyanCatda
 * @Description: 启动服务端
 * @FilePath: \Atsuko\StartServe.go
 */
package main

import (
	"context"
	"fmt"

	"github.com/nyancatda/Atsuko/internal/Flag"
	"github.com/nyancatda/Atsuko/internal/TCPComm"
)

func StartServe(Context context.Context) {
	go TCPComm.StartServe(Flag.Flag.ListenPort, Context, MessageChan, func(Msg string) {
		// 接收消息回调
		fmt.Println("\n" + Msg)

		// 打印命令提示符
		fmt.Print("\r>")
	})
}
