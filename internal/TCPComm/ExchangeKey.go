/*
 * @Author: NyanCatda
 * @Date: 2022-11-23 17:12:24
 * @LastEditTime: 2022-11-23 20:11:23
 * @LastEditors: NyanCatda
 * @Description: 交互密钥
 * @FilePath: \Atsuko\internal\TCPComm\ExchangeKey.go
 */
package TCPComm

import (
	"fmt"
	"net"
	"strings"
	"sync"

	"github.com/nyancatda/Atsuko/internal/MessageProcessing"
	"github.com/nyancatda/Atsuko/internal/SecretKey"
)

var KeysDir = "./keys/"                              // 密钥文件目录
var PersonKeyFile = KeysDir + "%s/%s/public_key.pem" // Person密钥文件路径

/**
 * @description: 交换密钥
 * @param {net.Conn} Conn 连接
 * @return {*}
 */
func ExchangeKey(Conn net.Conn) {
	// 交换密钥
	fmt.Println("正在与对方交换密钥...")

	var WaitGroup sync.WaitGroup

	// 发送自己的公钥
	WaitGroup.Add(1)
	go func() {
		Key, err := SecretKey.ReadPublicKey("./public_key.pem")
		if err != nil {
			fmt.Println("交换密钥失败: ", err)
			return
		}
		PublicKey, err := Key.PEMEncodePublicKey()
		if err != nil {
			fmt.Println("交换密钥失败: ", err)
			return
		}
		Conn.Write(PublicKey)

		WaitGroup.Done()
	}()

	// 接收对方的公钥
	WaitGroup.Add(1)
	go func() {
		Tmp := make([]byte, 4096)
		for {
			// 接收消息
			_, err := Conn.Read(Tmp[:])
			if err != nil {
				fmt.Println("链接已断开: ", err)
				ConnectionStatus = false
				break
			}

			if string(Tmp[:]) != "" {
				// 解析密钥
				Key, err := SecretKey.PEMDecodePublicKey(Tmp[:])
				if err != nil {
					fmt.Println("交换密钥失败: ", err)
					break
				}
				// 保存密钥
				AddressArray := strings.Split(Conn.RemoteAddr().String(), ":")
				err = Key.SavePublicKey(fmt.Sprintf(PersonKeyFile, AddressArray[0], AddressArray[1]))
				if err != nil {
					fmt.Println("交换密钥失败: ", err)
					break
				}

				break
			} else {
				fmt.Println("交换密钥失败")
				break
			}
		}

		WaitGroup.Done()
	}()

	// 等待密钥交换完成
	WaitGroup.Wait()

	// 将对方公钥写入MessageProcessing
	Personkey, err := SecretKey.ReadPublicKey(fmt.Sprintf(PersonKeyFile, strings.Split(Conn.RemoteAddr().String(), ":")[0], strings.Split(Conn.RemoteAddr().String(), ":")[1]))
	if err != nil {
		fmt.Println("交换密钥失败: ", err)
		return
	}
	MessageProcessing.Personkey = Personkey

	fmt.Println("密钥交换完成")
}
