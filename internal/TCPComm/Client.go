/*
 * @Author: NyanCatda
 * @Date: 2022-11-22 20:06:37
 * @LastEditTime: 2022-11-23 17:29:59
 * @LastEditors: NyanCatda
 * @Description: TCP客户端
 * @FilePath: \Atsuko\internal\TCPComm\Client.go
 */
package TCPComm

import (
	"fmt"
	"net"
)

/**
 * @description: 开启TCP客户端
 * @param {string} Address 连接地址
 * @param {func} Call 回调函数
 * @return {error} 错误
 */
func StartClient(Address string, MessageChan chan string, Call func(string, net.Conn)) {
	// 链接服务端
	fmt.Println(fmt.Sprintf("正在尝试与 %s 建立连接...", Address))
	Conn, err := net.Dial("tcp", Address)
	if err != nil {
		fmt.Println("无法建立连接: ", err)
		ConnectionStatus = false
		return
	}
	fmt.Println("连接建立成功")

	// 交换密钥
	ExchangeKey(Conn)

	ConnectionStatus = true

	go ReadProcess(Conn, Call)
	go WriteProcess(Conn, MessageChan)
}
