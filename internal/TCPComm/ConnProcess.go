/*
 * @Author: NyanCatda
 * @Date: 2022-11-23 14:01:05
 * @LastEditTime: 2022-11-23 19:49:08
 * @LastEditors: NyanCatda
 * @Description: 链接处理
 * @FilePath: \Atsuko\internal\TCPComm\ConnProcess.go
 */
package TCPComm

import (
	"fmt"
	"net"
)

var ConnectionStatus bool = false // 连接状态标识

/**
 * @description: 连接读取处理
 * @param {net.Conn} Conn 链接
 * @param {func} Call 回调函数
 * @param {func} ErrorCall 错误回调函数
 * @return {*}
 */
func ReadProcess(Conn net.Conn, Call func(string, net.Conn)) {
	Tmp := make([]byte, 4096)
	for {
		// 接收消息
		MesageLen, err := Conn.Read(Tmp[:])
		if err != nil {
			fmt.Println("链接已断开: ", err)
			ConnectionStatus = false
			break
		}

		// 执行回调函数
		if string(Tmp[:MesageLen]) != "" {
			Call(string(Tmp[:MesageLen]), Conn)
		}

		// 清空缓存
		Tmp = make([]byte, 4096)
	}
}

/**
 * @description: 连接写入处理
 * @param {net.Conn} Conn 链接
 * @param {chan string} MessageChan 消息通道
 * @return {*}
 */
func WriteProcess(Conn net.Conn, MessageChan chan string) {
	for {
		// 从消息通道获取消息
		Message := <-MessageChan

		// 发送消息
		if Message != "" {
			Conn.Write([]byte(Message))
		}
	}
}
