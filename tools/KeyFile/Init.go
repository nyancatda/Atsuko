/*
 * @Author: NyanCatda
 * @Date: 2022-11-23 16:55:27
 * @LastEditTime: 2022-11-23 17:05:01
 * @LastEditors: NyanCatda
 * @Description:
 * @FilePath: \Atsuko\tools\KeyFile\Init.go
 */
package KeyFile

import (
	"github.com/nyancatda/Atsuko/internal/SecretKey"
	"github.com/nyancatda/Atsuko/tools/File"
)

/**
 * @description: 初始化密钥文件
 * @return {error} 错误信息
 */
func Init() error {
	// 判断根目录是否存在密钥文件
	if File.Exists("./private_key.pem") && File.Exists("./public_key.pem") {
		// 如果存在，则直接返回
		return nil
	}

	// 如果不存在，则创建密钥文件
	Key, err := SecretKey.CreateKey()
	if err != nil {
		return err
	}
	err = Key.SaveKey("./")
	if err != nil {
		return err
	}

	return nil
}
