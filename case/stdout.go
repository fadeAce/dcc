package _case

import (
	"fmt"
	"io"
	"os"
)

func Output() {
	file, err := os.Open("./data/test.log")
	if err != nil {
		file, err = os.Create("./data/test.log")
		if err != nil {
			fmt.Println("err", err)
		}
	}
	n, _ := file.Seek(0, io.SeekEnd)
	_, err = file.WriteAt([]byte("find here"), n)
}
