package structure

type File struct {
	Name         string
	IsDir        bool
	Size         string
	ModifiedTime string
}
type DirEntry struct {
	Pointer string
	Cur     string
	Files   []File
}

var MyDirEntry = DirEntry{}
