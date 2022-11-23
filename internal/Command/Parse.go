/*
 * @Author: NyanCatda
 * @Date: 2022-07-19 16:32:52
 * @LastEditTime: 2022-11-23 14:28:44
 * @LastEditors: NyanCatda
 * @Description: 解析命令参数模块
 * @FilePath: \Atsuko\internal\Command\Parse.go
 */
package Command

import "strings"

/**
 * @description: 解析命令参数
 * @param {string} Command 命令字符串
 * @return {string} 命令
 * @return {[]string} 参数
 */
func Parse(Command string) (string, []string) {
	// 按照空格分割命令
	CommandArray := strings.Fields(Command)
	// 处理带冒号的命令参数
	CommandArray = ColonCommand(CommandArray)

	return CommandArray[0], CommandArray[1:]
}

/**
 * @description: 处理带冒号的命令参数
 * @param {[]string} CommandArray 命令内容
 * @return {[]string} 命令内容
 */
func ColonCommand(CommandArray []string) []string {
	var ColonCommandArray []string
	var ColonCommand string
	for _, Value := range CommandArray {
		if ColonCommand != "" {
			// 如果ColonCommand不为空，表示上个参数带有冒号，直接追加
			if strings.Contains(Value, `"`) {
				// 如果这个参数带冒号，则表示是参数结尾，追加进参数后清理ColonCommand
				ColonCommand = ColonCommand + " " + Value
				// 去除首尾冒号
				ColonCommand = strings.Trim(ColonCommand, `"`)
				ColonCommandArray = append(ColonCommandArray, ColonCommand)
				// 清空ColonCommand
				ColonCommand = ""
			} else {
				// 如果不带冒号，则带空格追加到ColonCommand
				ColonCommand = ColonCommand + " " + Value
			}
		} else {
			// 如果ColonCommand为空，表示上个参数不带有冒号，判断是否带有冒号
			if strings.Contains(Value, `"`) {
				ColonCommand = ColonCommand + Value
			} else {
				ColonCommandArray = append(ColonCommandArray, Value)
			}
		}
	}
	return ColonCommandArray
}
