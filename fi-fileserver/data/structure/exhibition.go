package structure

import (
	"log"
	"os"
	"strconv"
)

type File struct {
	Name         string
	IsDir        bool
	Size         string
	ModifiedTime string
}
type DirEntry struct {
	Pointer string //物理地址
	Cur     string //显示地址
	Files   []File
}

var MyDirEntry = DirEntry{}

func bmgt(int642 int64) string {
	switch {
	default:
		return ""
	case int642 < int64(1024):
		return strconv.FormatInt(int64(int642), 10) + "B"
	case int642 < int64(1024*1024):
		return strconv.FormatInt(int64(int642/1024), 10) + "KB"
	case int642 < int64(1024*1024*1024):
		return strconv.FormatInt(int64(int642/1024/1024), 10) + "MB"
	case int642 < int64(1024*1024*1024*1024):
		return strconv.FormatInt(int64(int642/1024/1024/1024), 10) + "GB"
	}
}
func (f File) UpdateDir(poi string, cur string, name string) DirEntry {
	log.Println("from exhibition.UpdateDir")
	poi += name + "/"
	cur += name + "/"
	log.Println("poi string,cur string,name string", poi, cur, name)
	dir, err := os.ReadDir(poi)
	if err != nil {
		log.Fatal(err)
	}
	direntry := DirEntry{}
	direntry.Cur = cur
	direntry.Pointer = poi
	for _, v := range dir {
		t, _ := v.Info()
		file := File{v.Name(), v.IsDir(), bmgt(t.Size()), t.ModTime().Format("2006年01月02日 15:04:05")}
		direntry.Files = append(direntry.Files, file)
	}
	log.Println("readdir:", direntry)

	log.Println("from exhibition.UpdateDir [end]")
	return direntry
}
