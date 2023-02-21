package gozip

import (
	"fmt"
	"testing"
)

func TestReadDirs(t *testing.T) {
	dir := "C:\\Users\\Administrator\\Desktop\\test_dir\\"

	filenames := ReadDirs(dir, dir)
	fmt.Println("filenames", len(filenames))
	fmt.Println(filenames[2].dirpath)
	fmt.Println(filenames[2].FileName)
	fmt.Println(filenames[2].GetAbsulutePath())
	fmt.Println(filenames[2].GetRelatedPath())
}
func Min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func TestZip(t *testing.T) {
	dir := "C:\\Users\\Administrator\\Desktop\\test_dir\\"
	zip := NewZip(dir)
	zip.ZIP("ziped")
	zip.CalaZipRate()
}
