<div align="center">

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ppdxzz/go-holiday?logo=go)
![Release](https://img.shields.io/github/v/release/ppdxzz/go-holiday?label=Release&color=blue)
![GitHub repo size](https://img.shields.io/github/repo-size/ppdxzz/go-holiday?label=Size&color=green)

</div>

### 📖 项目介绍
go-holiday 是一个Go版本的国内法定节假日API项目，它只需要接收一个日期值，就可以很方便的判断出这一天是不是节假日。


### 🔥 快速开始
```shell
go get -u github.com/ppdxzz/go-holiday
```
#### 判断节假日以及工作日
```go
package main

import (
    "fmt"
    "github.com/ppdxzz/go-holiday/holiday"
)

func main() {
    date := "2024-01-06"
    // IsHoliday 判断节假日 isInclude-true 包含双休
    isHoliday1, err := holiday.IsHoliday(date, true)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("日期%s是否为节假日(包含双休)：%t\n", date, isHoliday1)

    // IsHoliday 判断节假日 isInclude-false 不包含双休
    isHoliday2, err := holiday.IsHoliday(date, false)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("日期%s是否为节假日(不包含双休)：%t\n", date, isHoliday2)

    // IsWeekday 判断工作日，法定节假日调休也属于工作日
    isWeekday, err := holiday.IsWeekday(date)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("日期%s是否为工作日(包含调休)：%t\n", date, isWeekday)
}
```
_运行结果：_
```markdown
日期2024-01-06是否为节假日(包含双休)：true
日期2024-01-06是否为节假日(不包含双休)：false
日期2024-01-06是否为工作日(包含调休)：false
```


### ⚠️ 注意事项
由于接口发布年已然是2024年，故从24年开始维护节假日JSON数据，这也正是接口的使用限制之处，只能判断2024-01-01及之后的日期，之前的日期无法判断，之后的会持续更新数据源，望周知。


### 📃 开源协议
<a href="https://github.com/ppdxzz/go-holiday/blob/main/LICENSE">Apache License 2.0</a>