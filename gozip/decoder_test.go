package gozip

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestDecode(t *testing.T) {
	dir := "../ziped.gozip"
	originFile, err := ioutil.ReadFile("C:\\Users\\Administrator\\Desktop\\test_dir\\MP4test.mp4")

	filedata, err := ioutil.ReadFile(dir)
	if err != nil {
		panic(err)
	}

	decoder := NewDecoder(filedata)

	data := decoder.zipdata[0]
	fmt.Println("FileName :", data.FileName)
	for k, v := range data.CodeTable {
		fmt.Printf("| %v | %v |\n", k, v)
	}
	data.Decode()
	fmt.Printf("originFile Len %v originFile[last]: %v\n", len(originFile), originFile[len(originFile)-10:])

	fmt.Printf("data.UnZipedData Len %v data.UnZipedData[last]: %v\n", len(data.UnZipedData), data.UnZipedData[len(data.UnZipedData)-10:])
	// fmt.Printf("data.UnZipedData: %v\n", data.UnZipedData)
	// fmt.Printf("data.UnZipedData: %v\n", string(data.UnZipedData))
}
