/*
 * @Author: NyanCatda
 * @Date: 2022-07-19 15:58:02
 * @LastEditTime: 2022-11-23 14:37:57
 * @LastEditors: NyanCatda
 * @Description: 命令执行模块
 * @FilePath: \Atsuko\internal\Command\Execution.go
 */
package Command

import (
	"errors"
	"strings"
)

// 命令信息类型
type CommandInfo struct {
	Help     string
	Callback func(string)
}

// 命令列表
var CommandList map[string]CommandInfo

/**
 * @description: 执行命令
 * @param {string} Command 命令字符串
 * @return {error} 错误信息
 */
func Execution(Command string) error {
	// 去除换行符
	Command = strings.TrimSuffix(Command, "\n")

	// 分割命令
	CommandArray := strings.Fields(Command)
	// 判断命令是否为空
	if len(CommandArray) == 0 {
		return nil
	}

	// 循环匹配命令列表
	for CommandName, CommandInfo := range CommandList {
		if CommandName == CommandArray[0] {
			// 执行命令回调函数
			CommandInfo.Callback(Command)
			return nil
		}
	}

	// 如果没有匹配到任何命令则返回命令不存在
	return errors.New("Command not found")
}
