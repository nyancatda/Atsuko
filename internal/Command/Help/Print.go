/*
 * @Author: NyanCatda
 * @Date: 2022-07-19 17:06:36
 * @LastEditTime: 2022-11-23 14:31:17
 * @LastEditors: NyanCatda
 * @Description: 打印Help信息
 * @FilePath: \Atsuko\internal\Command\Help\Print.go
 */
package Help

import (
	"fmt"

	"github.com/nyancatda/Atsuko/internal/Command"
)

/**
 * @description: 打印Help信息
 * @param {*}
 * @return {*}
 */
func Print() {
	fmt.Println("---------------- Help ----------------")

	// 打印Help信息
	for Command, CommandInfo := range Command.CommandList {
		// 跳过help命令
		if Command == "help" || Command == "?" {
			continue
		}

		fmt.Println(Command + " " + CommandInfo.Help)
	}

	fmt.Println("--------------------------------------")
}
