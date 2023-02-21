package gozip

type FileInfo struct {
	OriginalDir  string
	OriginalPath string
	FileName     string
	dirpath      string
	data         []byte
}

func (f FileInfo) GetDate() []byte {
	return f.data
}
func (f FileInfo) GetFileName() string {
	return f.OriginalDir + "\\" + f.FileName
}
func (f FileInfo) GetAbsulutePath() string {
	return f.OriginalDir + "\\" + f.FileName
}
func (f FileInfo) GetRelatedPath() string {
	return f.dirpath[len(f.OriginalPath)+1:] + f.FileName
}
