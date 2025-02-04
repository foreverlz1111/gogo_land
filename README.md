# 安装
从[[官网下载页面](https://go.dev/dl/)]获取对应系统的二进制包，

解压包内文件到/usr/local/go/然后添加用户变量：

```
echo export PATH=$PATH:/usr/local/go/bin >> ~/.bashrc
```

重启电脑，然后尝试一下：

```
go version
```

运行代码：

```
go run main.go
```

**Go 语言不支持动态链接，因此编译时会将所有依赖编译进同一个二进制文件：**

```
go build -o main main.go
```

**交叉编译（示例）：**

|GOOS |GOARCH | 
|-|-|
|linux | amd64|
|linux | arm64|
|freebsd | amd64|
|freebsd | arm64|
|darwin | amd64|
|darwin | arm64|
|...|[更多](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63)|
 
```
GOOS=linux GOARCH=amd64 go build -o myprogram main.go
``` 

---

# 文件夹
- [go1](https://github.com/foreverlz1111/gogo_land/tree/main/go1)-
`
包、变量和函数
`
- [go2](https://github.com/foreverlz1111/gogo_land/tree/main/go2)-
`
流程控制语句：for、if、else、switch 和 defer
`
- [go3](https://github.com/foreverlz1111/gogo_land/tree/main/go3)-
`
更多类型：struct、slice 和映射
`
- [go4](https://github.com/foreverlz1111/gogo_land/tree/main/go4)-
`
方法和接口
`

- [go5](https://github.com/foreverlz1111/gogo_land/tree/main/go5)-
`
并发
`

- [go6](https://github.com/foreverlz1111/gogo_land/tree/main/go6)-
`
排序算法
`

- [leet](https://github.com/foreverlz1111/gogo_land/tree/main/leet)- code

- [sub-knowledge](https://github.com/foreverlz1111/gogo_land/tree/main/sub-knowledge)-
`
课外知识
`