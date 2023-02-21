package gozip

import (
	"log"
)

type ZipData struct {
	FileDir     string
	FileName    string
	CodeTable   map[string]byte
	Body        []byte
	UnZipedData []byte
}

func (z *ZipData) Decode() error {
	z.UnZipedData = []byte{}
	s := ""
	bodyLen := len(z.Body)
	lastLen := z.Body[bodyLen-1]
	for idx, code := range z.Body[:bodyLen-1] {
		begin := 7
		if idx == bodyLen-2 && lastLen == 1 {
			begin = 0
		} else if idx == bodyLen-2 && lastLen != 1 {
			begin = int(lastLen) - 1
		}

		for i := begin; i >= 0; i-- {
			if (code>>i)&1 == 1 {
				s += "1"
			} else {
				s += "0"
			}
			if b, ok := z.CodeTable[s]; ok {
				z.UnZipedData = append(z.UnZipedData, b)
				s = ""
			}
		}
	}
	if s != "" {
		log.Panicln("s != null\n lastLen:", lastLen)
	}

	return nil
}
