/*
 * @Author: NyanCatda
 * @Date: 2022-11-22 20:00:04
 * @LastEditTime: 2022-11-22 22:02:24
 * @LastEditors: NyanCatda
 * @Description: TCP服务端
 * @FilePath: \Atsuko\internal\TCPComm\Serve.go
 */
package TCPComm

import (
	"net"
	"strconv"
)

/**
 * @description: 启动TCP服务端
 * @param {int} Port 端口号
 * @param {func} Call 回调函数
 * @param {func} ErrorCall 错误回调函数
 * @return {*}
 */
func StartServe(Port int, MessageChan chan string, Call func(string), ErrorCall func(error)) {
	// 启动监听访问
	Listener, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(Port))
	if err != nil {
		ErrorCall(err)
		return
	}

	// 等待通信建立
	for {
		Conn, err := Listener.Accept()
		if err != nil {
			ErrorCall(err)
			continue
		}
		go ServerReadProcess(Conn, Call, ErrorCall)
		go ServerWriteProcess(Conn, MessageChan)
	}
}

/**
 * @description: 服务端读取处理
 * @param {net.Conn} Conn 连接
 * @param {func} ReceiveCall 回调函数
 * @return {*}
 */
func ServerReadProcess(Conn net.Conn, ReceiveCall func(string), ErrorCall func(error)) {
	Tmp := make([]byte, 4096)
	for {
		_, err := Conn.Read(Tmp[:])
		if err != nil {
			ErrorCall(err)
			break
		}

		// 执行回调函数
		if string(Tmp[:]) != "" {
			ReceiveCall(string(Tmp[:]))
		}

		// 清空缓存
		Tmp = make([]byte, 4096)
	}
}

/**
 * @description: 服务端写入处理
 * @param {net.Conn} Conn 连接
 * @param {chan string} MessageChan 消息通道
 * @return {*}
 */
func ServerWriteProcess(Conn net.Conn, MessageChan chan string) {
	for {
		// 从消息通道获取消息
		Message := <-MessageChan
		if Message != "" {
			Conn.Write([]byte(Message))
		}
	}
}
