package structure

import "time"

type Header struct {
	Title     string
	HintTitle string
	Date      string
	Upload    string
	Edit      string
}
type list struct {
	List string
}

var Mylist = list{"list1"}
var loc = time.Local
var Myheader = Header{
	Title:     "标题",
	HintTitle: "文件预览",
	Edit:      "编辑",
	Upload:    "上传",
	Date:      time.Date(2022, 11, 7, 0, 0, 0, 0, loc).Format("2006年01月02日"),
}
