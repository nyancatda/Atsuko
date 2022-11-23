/*
 * @Author: NyanCatda
 * @Date: 2022-11-23 16:19:15
 * @LastEditTime: 2022-11-23 16:53:44
 * @LastEditors: NyanCatda
 * @Description: 密钥文件操作
 * @FilePath: \Atsuko\internal\SecretKey\File.go
 */
package SecretKey

import (
	"os"
	"path/filepath"

	"github.com/nyancatda/Atsuko/tools/File"
)

/**
 * @description: 保存私钥到文件
 * @param {string} FilePath 文件路径
 * @return {error} 错误信息
 */
func (Key Key) SavePrivateKey(FilePath string) error {
	Path, _ := filepath.Split(FilePath)

	// 获取密钥PEM
	KeyPEM := Key.PEMEncodePrivateKey()

	// 创建文件夹
	File.MKDir(Path)

	// 保存文件
	FileWrite, err := File.NewFileReadWrite(FilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE)
	if err != nil {
		return err
	}
	defer FileWrite.Close()

	err = FileWrite.WriteTo(string(KeyPEM))
	if err != nil {
		return err
	}

	return nil
}

/**
 * @description: 保存公钥到文件
 * @param {string} FilePath 文件路径
 * @return {error} 错误信息
 */
func (Key Key) SavePublicKey(FilePath string) error {
	Path, _ := filepath.Split(FilePath)

	// 获取密钥PEM
	KeyPEM, err := Key.PEMEncodePublicKey()
	if err != nil {
		return err
	}

	// 创建文件夹
	File.MKDir(Path)

	// 保存文件
	FileWrite, err := File.NewFileReadWrite(FilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE)
	if err != nil {
		return err
	}
	defer FileWrite.Close()

	err = FileWrite.WriteTo(string(KeyPEM))
	if err != nil {
		return err
	}

	return nil
}

/**
 * @description: 储存密钥对到文件夹
 * @param {string} DirPath 文件夹路径
 * @return {error} 错误信息
 */
func (Key Key) SaveKey(DirPath string) error {
	err := Key.SavePrivateKey(DirPath + "/private_key.pem")
	if err != nil {
		return err
	}
	err = Key.SavePublicKey(DirPath + "/public_key.pem")
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: 读取私钥文件
 * @param {string} FilePath 文件路径
 * @return {Key} 密钥对
 * @return {error} 错误信息
 */
func ReadPrivateKey(FilePath string) (Key, error) {
	if !File.Exists(FilePath) {
		return Key{}, os.ErrNotExist
	}

	// 读取文件
	FileRead, err := File.NewFileReadWrite(FilePath, os.O_RDONLY)
	if err != nil {
		return Key{}, err
	}
	defer FileRead.Close()
	FileBody, err := FileRead.Read()
	if err != nil {
		return Key{}, err
	}

	// 解析PEM
	var Key Key
	Key, err = PEMDecodePrivateKey([]byte(FileBody))
	if err != nil {
		return Key, err
	}

	return Key, nil
}

/**
 * @description: 读取公钥文件
 * @param {string} FilePath 文件路径
 * @return {Key} 密钥对
 */
func ReadPublicKey(FilePath string) (Key, error) {
	if !File.Exists(FilePath) {
		return Key{}, os.ErrNotExist
	}

	// 读取文件
	FileRead, err := File.NewFileReadWrite(FilePath, os.O_RDONLY)
	if err != nil {
		return Key{}, err
	}
	defer FileRead.Close()
	FileBody, err := FileRead.Read()
	if err != nil {
		return Key{}, err
	}

	// 解析PEM
	var Key Key
	Key, err = PEMDecodePublicKey([]byte(FileBody))
	if err != nil {
		return Key, err
	}

	return Key, nil
}

/**
 * @description: 从文件夹读取private_key.pem与public_key.pem文件
 * @param {string} DirPath 文件夹路径
 * @return {Key} 密钥对
 * @return {error} 错误信息
 */
func ReadKey(DirPath string) (Key, error) {
	PrivateKey, err := ReadPrivateKey(DirPath + "/private_key.pem")
	if err != nil {
		return Key{}, err
	}
	PublicKey, err := ReadPublicKey(DirPath + "/public_key.pem")
	if err != nil {
		return Key{}, err
	}

	// 合并密钥对
	var Key Key
	Key.PrivateKey = PrivateKey.PrivateKey
	Key.PublicKey = PublicKey.PublicKey

	return Key, nil
}
