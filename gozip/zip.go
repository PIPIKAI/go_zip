package gozip

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
)

type Zip struct {
	CodedDatas []CodedData
}

// 递归便利文件夹下所有文件
func ReadDirs(orgPath string, dir string) []FileInfo {
	var fileinfo []fs.FileInfo
	if filestate, err := os.Stat(orgPath); err != nil {
		panic(err)
	} else {
		if !filestate.IsDir() {
			fileinfo = []fs.FileInfo{filestate}
			dir = "."
		} else {
			fileinfo, err = ioutil.ReadDir(dir)
			if err != nil {
				panic(err)
			}
		}
	}

	temp := []FileInfo{}
	for _, fs := range fileinfo {
		if fs.IsDir() {
			res := ReadDirs(orgPath, dir+"\\"+fs.Name())
			temp = append(temp, res...)
		} else {
			filedata, err := ioutil.ReadFile(dir + "\\" + fs.Name())
			if err != nil {
				panic(fmt.Sprintf("read file %v \n[ERR]: %v!", fs.Name(), err.Error()))
			}

			OriginalDir := dir[strings.LastIndex(orgPath, "\\")+1:]
			temp = append(temp, FileInfo{
				OriginalDir:  OriginalDir,
				OriginalPath: orgPath,
				FileName:     fs.Name(),
				dirpath:      dir + "\\",
				data:         filedata,
			})
		}
	}
	return temp
}

func NewZip(filePaths ...string) *Zip {

	codeddata := []CodedData{}
	// 这里考虑加多线程
	for _, filePath := range filePaths {
		fileinfos := ReadDirs(filePath, filePath)
		for _, fileinfo := range fileinfos {
			tempData := CodedData{
				OriginData: fileinfo,
			}
			tempData.EncodeBody()
			tempData.EncodeHead()

			codeddata = append(codeddata, tempData)
		}

	}

	return &Zip{
		CodedDatas: codeddata,
	}

}

func (z Zip) ZIP(dir string) {
	zipedData := []byte{}
	for _, v := range z.CodedDatas {
		data := append(v.head, v.body...)

		zipedData = append(zipedData, data...)
	}
	if err := ioutil.WriteFile(dir+".gozip", zipedData, 0666); err != nil {
		panic(err.Error())
	}
}
func (z Zip) CalaZipRate() {
	var originalLen int64
	var codeedLen int64
	for _, codeddata := range z.CodedDatas {
		originalLen += int64(len(codeddata.OriginData.data))
		codeedLen += int64(len(codeddata.body) + len(codeddata.head))
	}
	fmt.Printf("coded / origin  = %f  \n", (float32(codeedLen) / float32(originalLen*1.0)))
}
