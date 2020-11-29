# task

`go run .`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>      
* [func AppendFile(path string, content string) error](#AppendFile)
* [func CheckOrCreateFile(filepath string) error](#CheckOrCreateFile)
* [func RoutineWork(c *Config, msg chan&lt;- string)](#RoutineWork)
* [type Config](#Config)



## <a name="AppendFile">func</a> AppendFile
``` go
func AppendFile(path string, content string) error
```
AppendFile 追加文件内容



## <a name="CheckOrCreateFile">func</a> CheckOrCreateFile
``` go
func CheckOrCreateFile(filepath string) error
```
CheckOrCreateFile 检查文件是否存在，不存在则创建



## <a name="RoutineWork">func</a> RoutineWork
``` go
func RoutineWork(c *Config, msg chan<- string)
```
RoutineWork 开始协程，进行 DOS 攻击




## <a name="Config">type</a> Config
``` go
type Config struct {
    Log      string
    URL      string
    Query    string
    Routine  int
    ReqTimes int
}

```
Config DOS 配置内容