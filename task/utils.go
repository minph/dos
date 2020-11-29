package task

import (
	"fmt"
	"os"
)

// AppendFile 追加文件内容
func AppendFile(path string, content string) error {
	// 追加方式打开文件
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println(err)
		return err
	}

	// 写入数据
	_, err = file.Write([]byte(content))

	if err != nil {
		fmt.Println(err)
		return err
	}

	// 关闭文件
	err = file.Close()

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// CheckOrCreateFile 检查文件是否存在，不存在则创建
func CheckOrCreateFile(filepath string) error {
	_, err := os.Stat(filepath)

	// 文件存在直接结束
	if err == nil {
		return nil
	}

	// 文件不存在，则创建文件
	if os.IsNotExist(err) {
		_, err = os.Create(filepath)
		return err
	}

	return err
}
