package gozip

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestEncodeBody(t *testing.T) {
	dataDir := "C:\\Users\\Administrator\\Desktop\\test_dir\\MP4test.mp4"
	data, err := ioutil.ReadFile(dataDir)
	if err != nil {
		panic(err)
	}
	temp := FileInfo{
		OriginalDir:  "TimelineDemo",
		OriginalPath: dataDir,
		FileName:     "test.txt",
		data:         data,
	}
	tempData := CodedData{
		OriginData: temp,
	}
	tempData.EncodeBody()
	tempData.EncodeHead()

	for k, v := range tempData.codeTable {
		fmt.Printf("|%4v|%4v|\n", k, v)
	}
	// fmt.Println(string(tempData.OriginData.data))
	fmt.Println(tempData.OriginData.data)
	for _, v := range tempData.body[4:] {
		fmt.Printf("%b|", v)
	}
}
