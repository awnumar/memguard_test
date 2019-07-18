package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
)

var key = []byte{0xd8, 0x69, 0xb4, 0x65, 0x42, 0x99, 0xc0, 0x65, 0xce, 0x24, 0x9c, 0x2f, 0x27, 0xcf, 0xb8, 0x28, 0xdc, 0xc, 0xa9, 0xdd, 0x23, 0x30, 0x35, 0xfe, 0x91, 0x90, 0xc2, 0x4, 0x8f, 0x8e, 0x12, 0x7c}

func main() {
	var dumps []string

	filepath.Walk("./memdump", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && path[len(path)-5:] == ".dump" {
			dumps = append(dumps, path)
		}
		return nil
	})

	for _, dump := range dumps {
		f, err := os.Open(dump)
		if err != nil {
			panic(err)
		}

		data, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}

		f.Close()

		for i := 0; i <= cap(data)-32; i++ {
			if bytes.Equal(data[i:i+32], key) {
				panic("found")
			}
		}
	}
}
