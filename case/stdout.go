package _case

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Output() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// case using golang image
	if dir == "/go" {
		dir = "/usr/local/bin"
	}
	// case using ubuntu image
	if dir == "/" {
		dir = "/usr/local/bin"
	}
	file, err := os.OpenFile(dir+"/data/test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("err", err)
		file, err = os.Create(dir + "/data/test.log")
		if err != nil {
			fmt.Println("err", err)
		}
	}
	n, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.WriteAt([]byte("find here\n"), n)
	if err != nil {
		fmt.Println(err)
	}
}
