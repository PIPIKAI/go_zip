package gozip

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type UnZip struct {
	zipdata   []ZipData
	orginData []byte
}

func byte232Decode(arr []byte) int {
	res := 0
	offset := 1
	for i := 3; i >= 0; i-- {
		res = int(arr[i])*offset + res
		offset *= 256
	}
	return res
}
func NewDecoder(originData []byte) *UnZip {
	decoder := &UnZip{
		orginData: originData,
	}

	decoder.UnMashall()

	return decoder
}

func (d *UnZip) UnMashall() {
	zipdatas := []ZipData{}

	copyData := make([]byte, len(d.orginData))
	copy(copyData, d.orginData)

	for {
		var newZipData ZipData
		fileNameLen := byte232Decode(copyData[:4])
		fileName := string(copyData[4:fileNameLen])
		idx := strings.LastIndex(fileName, "\\") + 1
		newZipData.FileName = fileName[idx:]
		newZipData.FileDir = fileName[:idx]
		codeTables := make(map[string]byte)
		copyData = copyData[fileNameLen:]
		tableLen := byte232Decode(copyData[:4])

		for j := 0; j < tableLen-4; {
			v := copyData[4+j]
			klen := int(copyData[4+j+1])
			k := string(copyData[4+j+2 : 4+j+2+klen])
			codeTables[k] = v
			j += klen + 2
		}
		newZipData.CodeTable = codeTables
		copyData = copyData[tableLen:]
		bodyLen := byte232Decode(copyData[:4])
		body := copyData[4:bodyLen]
		newZipData.Body = body
		copyData = copyData[bodyLen:]
		zipdatas = append(zipdatas, newZipData)
		// log.Println("FileName:", newZipData.FileName, "FileDir:", newZipData.FileDir, "FileSize:", len(newZipData.Body)/1024, "k")

		if len(copyData) <= 0 {
			break
		}
	}
	d.zipdata = zipdatas
}
func (d *UnZip) UnZip(dir string) {
	for _, v := range d.zipdata {
		log.Println("FileName:", v.FileName, "FileDir:", v.FileDir, "FileSize:", len(v.Body)/1024, "k")

		v.Decode()

		var fileDir string
		if dir != "" {
			fileDir = dir + "\\" + v.FileDir
		} else {
			fileDir = v.FileDir
		}
		os.MkdirAll(fileDir, 0666)
		if err := ioutil.WriteFile(fileDir+v.FileName, v.UnZipedData, 0666); err != nil {
			panic(err.Error())
		}
	}
}
