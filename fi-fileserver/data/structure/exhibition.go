package structure

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
