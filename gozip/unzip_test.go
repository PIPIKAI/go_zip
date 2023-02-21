package gozip

import (
	"io/ioutil"
	"testing"
)

func TestUnMashall(t *testing.T) {
	dir := "../ziped.gozip"
	filedata, err := ioutil.ReadFile(dir)
	if err != nil {
		panic(err)
	}

	NewDecoder(filedata)

}

func TestUnZip(t *testing.T) {
	dir := "../ziped.gozip"
	filedata, err := ioutil.ReadFile(dir)
	if err != nil {
		panic(err)
	}

	data := NewDecoder(filedata)
	data.UnZip("unziped")
}
