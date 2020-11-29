package main

import (
	"fmt"
	"strconv"

	"github.com/minph/conf"
)

// AppConfig 读取配置文件数据
var AppConfig = conf.New("app.conf")

// 读取数据并转 int 类型
func configInt(groupName, itemName string) int {

	// 获取字符串类型数据
	data := AppConfig.GetValue(groupName, itemName)
	value, err := strconv.Atoi(data)

	if err != nil {
		fmt.Println(err)
		return 0
	}

	return value
}
