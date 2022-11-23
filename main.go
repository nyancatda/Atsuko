/*
 * @Author: NyanCatda
 * @Date: 2022-11-21 19:54:01
 * @LastEditTime: 2022-11-23 15:32:39
 * @LastEditors: NyanCatda
 * @Description: 主文件
 * @FilePath: \Atsuko\main.go
 */
package main

import (
	"context"

	"github.com/nyancatda/Atsuko/internal/Flag"
)

var MessageChan = make(chan string) // 创建消息通道

func main() {
	// 初始化参数
	Flag.Init()

	// 启动服务端
	Context, CancelServer := context.WithCancel(context.Background()) // 创建上下文
	StartServe(Context)

	// 注册命令
	CommandRegister(CancelServer)

	// 启动控制台
	StartConsole()
}
