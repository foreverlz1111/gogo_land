package structure

type Footer struct {
	Title string
	Todo  []string
}

var Myfooter = Footer{
	Title: "footer",
	Todo: []string{
		"上一页 ?",
		"日志记录格式加入日期",
		"标题显示当前目录",
		"编辑文件目录",
		"下载文件按钮-浏览器打不开的默认下载",
		"点击文件进行浏览-浏览器可读的pdf、txt文件默认浏览",
	},
}
