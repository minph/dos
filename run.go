package main

import (
	"fmt"

	"github.com/minph/dos/task"
)

// Run DOS 任务调配
func Run() {

	// 新建数据通道

	msg := make(chan string)

	// 获取配置信息
	c := &task.Config{
		Log:      AppConfig.GetValue("app", "log"),
		URL:      AppConfig.GetValue("app", "url"),
		Query:    AppConfig.GetValue("app", "query"),
		Routine:  configInt("app", "routine"),
		ReqTimes: configInt("app", "reqTimes"),
	}

	fmt.Println("目标地址 = ", c.URL)
	fmt.Println("开启协程 = ", c.Routine)
	fmt.Println("协程循环 = ", c.ReqTimes)
	fmt.Println("日志文件 = ", c.Log)

	// 检查日志文件
	if len(c.Log) > 1 {

		// 检查存在，否则创建
		err := task.CheckOrCreateFile(c.Log)
		if err != nil {
			fmt.Println(err)
		}
	}

	// 设置协程计数器
	counter := c.Routine

	// 开启协程
	go task.RoutineWork(c, msg)

	for {

		message, ok := <-msg

		if !ok {
			continue
		}

		// 判断协程发送结束信号
		if len(message) == 0 {

			// 计数器自减
			counter--

			if counter == 0 {
				// 判断协程是否全部结束
				fmt.Println("All routines are done")

				break
			}
		} else {
			// 打印消息
			fmt.Println(message)
		}

	}

	// 暂停退出
	fmt.Scanln()

}

func main() {
	Run()
}
