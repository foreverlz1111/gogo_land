从[[官网](https://golang.org )]下载tar.gz包根据步骤解压安装

解压包到usr/local然后添加用户变量：
```
echo export PATH=$PATH:/usr/local/go/bin >>~/.bashrc
```
重启电脑，然后尝试一下：
```
go version
```
运行代码
```
go run main.go
```
**翻译自[go官网](https://tour.go-zh.org)**
