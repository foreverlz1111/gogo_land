package structure

type Dir struct {
}
type DirEntry struct {
	Pointer string
	Dirwalk []Dir
}
