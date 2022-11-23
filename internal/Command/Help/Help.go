/*
 * @Author: NyanCatda
 * @Date: 2022-11-23 14:30:53
 * @LastEditTime: 2022-11-23 14:32:21
 * @LastEditors: NyanCatda
 * @Description: 帮助命令模块
 * @FilePath: \Atsuko\internal\Command\Help\Help.go
 */
package Help

import "github.com/nyancatda/Atsuko/internal/Command"

/**
 * @description: 注册Help命令
 * @return {*}
 */
func Register() {
	// 注册Help命令
	Command.Add("help", "查看帮助", func(string) {
		Print()
	})
	Command.Add("?", "查看帮助", func(string) {
		Print()
	})
}
