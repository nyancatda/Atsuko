/*
 * @Author: NyanCatda
 * @Date: 2022-11-23 14:11:17
 * @LastEditTime: 2022-11-23 15:34:50
 * @LastEditors: NyanCatda
 * @Description: 参数读取
 * @FilePath: \Atsuko\internal\Flag\Flag.go
 */
package Flag

import "flag"

type Flags struct {
	ListenPort int
	Connect    string
}

var Flag Flags

/**
 * @description: 初始化参数
 * @return {error} 错误信息
 */
func Init() error {
	// 获取参数
	ListenPort := flag.Int("listen_port", 7788, "服务端模式监听端口")
	Connect := flag.String("connect", "", "客户端模式连接地址")
	flag.Parse()

	// 参数写入变量
	Flag.ListenPort = *ListenPort
	Flag.Connect = *Connect

	return nil
}
