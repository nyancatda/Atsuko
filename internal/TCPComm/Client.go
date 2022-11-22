/*
 * @Author: NyanCatda
 * @Date: 2022-11-22 20:06:37
 * @LastEditTime: 2022-11-22 22:02:36
 * @LastEditors: NyanCatda
 * @Description: TCP客户端
 * @FilePath: \Atsuko\internal\TCPComm\Client.go
 */
package TCPComm

import (
	"net"
)

/**
 * @description: 开启TCP客户端
 * @param {int} Port 端口号
 * @param {func} Call 回调函数
 * @return {error} 错误
 */
func StartClient(Address string, MessageChan chan string, Call func(string), ErrorCall func(error)) {
	// 链接服务端
	Conn, err := net.Dial("tcp", Address)
	if err != nil {
		ErrorCall(err)
		return
	}

	go ClientReadProcess(Conn, Call, ErrorCall)
	go ClientWriteProcess(Conn, MessageChan)
}

/**
 * @description: 客户端读取处理
 * @param {net.Conn} Conn 链接
 * @param {func} Call 回调函数
 * @param {func} ErrorCall 错误回调函数
 * @return {*}
 */
func ClientReadProcess(Conn net.Conn, Call func(string), ErrorCall func(error)) {
	Tmp := make([]byte, 4096)
	for {
		// 接收消息
		_, err := Conn.Read(Tmp[:])
		if err != nil {
			ErrorCall(err)
			break
		}

		// 执行回调函数
		if string(Tmp[:]) != "" {
			Call(string(Tmp[:]))
		}

		// 清空缓存
		Tmp = make([]byte, 4096)
	}
}

/**
 * @description: 客户端写入处理
 * @param {net.Conn} Conn 链接
 * @param {chan string} MessageChan 消息通道
 * @return {*}
 */
func ClientWriteProcess(Conn net.Conn, MessageChan chan string) {
	for {
		// 从消息通道获取消息
		Message := <-MessageChan

		// 发送消息
		if Message != "" {
			Conn.Write([]byte(Message))
		}
	}
}
