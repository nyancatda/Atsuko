/*
 * @Author: NyanCatda
 * @Date: 2022-11-21 19:54:01
 * @LastEditTime: 2022-11-23 17:05:24
 * @LastEditors: NyanCatda
 * @Description: 主文件
 * @FilePath: \Atsuko\main.go
 */
package main

import (
	"context"
	"fmt"

	"github.com/nyancatda/Atsuko/internal/Flag"
	"github.com/nyancatda/Atsuko/tools/KeyFile"
)

var MessageChan = make(chan string) // 创建消息通道

func main() {
	// 初始化参数
	if err := Flag.Init(); err != nil {
		fmt.Println(err)
		return
	}

	// 初始密钥文件
	if err := KeyFile.Init(); err != nil {
		fmt.Println(err)
		return
	}

	var CancelServer context.CancelFunc
	if Flag.Flag.Connect != "" {
		// 如果存在连接参数，则直接启动客户端模式
		Connect(Flag.Flag.Connect)
	} else {
		// 启动服务端
		Context, Cancel := context.WithCancel(context.Background()) // 创建上下文
		CancelServer = Cancel
		StartServe(Context)
	}

	// 注册命令
	CommandRegister(CancelServer)

	// 启动控制台
	StartConsole()
}
