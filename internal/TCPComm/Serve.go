/*
 * @Author: NyanCatda
 * @Date: 2022-11-22 20:00:04
 * @LastEditTime: 2022-11-23 15:55:26
 * @LastEditors: NyanCatda
 * @Description: TCP服务端
 * @FilePath: \Atsuko\internal\TCPComm\Serve.go
 */
package TCPComm

import (
	"context"
	"fmt"
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
func StartServe(Port int, Context context.Context, MessageChan chan string, Call func(string, net.Conn)) {
	// 启动监听访问
	Listener, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(Port))
	if err != nil {
		fmt.Println("无法启动服务:", err)
		return
	}
	fmt.Println("服务已启动，正在等待连接...")

	// 等待通信建立
	for {
		Conn, err := Listener.Accept()
		if err != nil {
			fmt.Println("无法建立连接: ", err)
			ConnectionStatus = false
			continue
		}

		select {
		case <-Context.Done():
			// 如果收到退出信号，关闭连接和监听
			Conn.Close()
			Listener.Close()
			return
		default:
			fmt.Println(fmt.Sprintf("成功与 %s 建立连接", Conn.RemoteAddr().String()))
			ConnectionStatus = true

			go ReadProcess(Conn, Call)
			go WriteProcess(Conn, MessageChan)
		}
	}
}
