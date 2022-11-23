/*
 * @Author: NyanCatda
 * @Date: 2022-11-23 15:27:01
 * @LastEditTime: 2022-11-23 15:38:01
 * @LastEditors: NyanCatda
 * @Description: 控制台处理
 * @FilePath: \Atsuko\Console.go
 */
package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/nyancatda/Atsuko/internal/Command"
	"github.com/nyancatda/Atsuko/internal/Command/Help"
)

/**
 * @description: 命令注册
 * @param {context.CancelFunc} CancelServer 服务端关闭函数
 * @return {*}
 */
func CommandRegister(CancelServer context.CancelFunc) {
	Command.Add("connect", "连接到对方客户端", func(CommandStr string) {
		// 关闭服务端
		CancelServer()

		// 解析命令参数
		_, Parameter := Command.Parse(CommandStr)
		if len(Parameter) != 1 {
			fmt.Println("命令格式错误，应为：connect [Host:Port]")
			return
		}

		// 连接到服务端
		Connect(Parameter[0])
	})

	Help.Register()
}

/**
 * @description: 启动控制台
 * @return {*}
 */
func StartConsole() {
	Reader := bufio.NewReader(os.Stdin)
	// 循环处理输入
	for {
		// 处理命令提示符输入
		Scanner := bufio.NewScanner(os.Stdin)
		if string(Scanner.Bytes()) == "" {
			fmt.Print("\r")
		} else {
			fmt.Print("\n\r")
		}

		// 获取输入
		CmdString, err := Reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		// 执行命令
		err = Command.Execution(CmdString)
		if err != nil {
			fmt.Println(err)
		}
	}
}
