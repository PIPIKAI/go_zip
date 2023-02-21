package gozip

import (
	"log"

	"github.com/pipikai/go_zip/haffman"
)

type CodedData struct {
	OriginData  FileInfo
	codeTable   map[byte]string
	head        []byte
	body        []byte
	weightTable map[byte]int64
}

func byte232Encode(numb int) []byte {
	res := make([]byte, 4)
	for i := 3; i >= 0; i-- {
		res[i] = byte(numb % 256)
		numb /= 256
	}
	return res
}

func (c *CodedData) EncodeHead() {
	res := []byte{}
	// 4 * 8
	// 构造一种数据结构 使得 将 32位的数字转化为byte 转化为 string
	// [  32(name_idx_end)  32 :table_idx_end  32:body_idx_end ]
	// table
	// key1 32(v1len) v1 , key2 32(v2len) v2 , ...
	//
	log.Println("FileName:", c.OriginData.GetRelatedPath(), "FileDir:", c.OriginData.GetFileName(), "FileSize:", len(c.OriginData.data)/1024, "k")

	fileName := c.OriginData.GetRelatedPath()
	table_idx_end := len(fileName) + 4

	res = append(res, byte232Encode(table_idx_end)...)
	res = append(res, []byte(fileName)...)

	tableCoded := []byte{}
	for k, v := range c.codeTable {
		tableCoded = append(tableCoded, k)
		if len(v) > 256 {
			panic("v > 256")
		}
		tableCoded = append(tableCoded, byte(len(v)))
		tableCoded = append(tableCoded, []byte(v)...)
	}
	res = append(res, byte232Encode(4+len(tableCoded))...)
	res = append(res, tableCoded...)
	c.head = res
}

func (c *CodedData) EncodeBody() {
	haffman := haffman.NewHaffmanCode(c.OriginData.GetDate())

	c.codeTable = haffman.GetCodeTable()
	c.weightTable = haffman.GetWeights()

	codeddatas := []byte{}

	cnt := 0
	now_v := 0
	for _, bt := range c.OriginData.data {
		k := c.codeTable[bt]
		for _, c := range k {
			now_v = now_v*2 + int(c-'0')
			if (cnt+1)%8 == 0 {
				codeddatas = append(codeddatas, byte(now_v))
				now_v = 0
			}
			cnt = (cnt + 1) % 8
		}
	}
	// if now_v != 0 {
	codeddatas = append(codeddatas, byte(now_v))
	// }
	codeddatas = append(codeddatas, byte(cnt))

	byteLens := byte232Encode(4 + len(codeddatas))

	c.body = append(byteLens, codeddatas...)
}
