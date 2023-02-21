package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	gozip "github.com/pipikai/go_zip/gozip"
)

var (
	mode     = flag.Bool("m", false, "默认为压缩")
	srcDir   = flag.String("s", "", "待压缩/解压的文件/夹")
	detDir   = flag.String("d", *srcDir, "压缩/解压的文件名/目录")
	calcRate = flag.Bool("c", false, "计算压缩比")
)

func main() {

	flag.Parse()

	if *srcDir == "" {
		fmt.Println("SrcDir Required")
		return
	}

	*srcDir = strings.Replace(*srcDir, "\\", "/", -1)
	*srcDir = strings.TrimRight(*srcDir, "/")

	fmt.Println("srcDir", *srcDir)

	if !*mode {
		if *detDir == "" {
			ss := strings.Split(*srcDir, "/")
			*detDir = ss[len(ss)-1]
		}
		fmt.Println("detDir", *detDir+".gozip")

		zip := gozip.NewZip(*srcDir)
		zip.ZIP(*detDir)

		if *calcRate {
			zip.CalaZipRate()
		}
	} else {
		if *detDir == "" {
			ss := strings.Split(*srcDir, ".")
			*detDir = ss[len(ss)-2]
			*detDir = strings.TrimLeft(*detDir, "/")
		}
		fmt.Println("detDir", *detDir)
		filedata, err := ioutil.ReadFile(*srcDir)
		if err != nil {
			panic(err)
		}

		data := gozip.NewDecoder(filedata)
		data.UnZip(*detDir)
	}

}
