package _case

import (
	"fmt"
	"io"
	"os"
)

func Output() {
	file, err := os.OpenFile("./data/test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("err", err)
		file, err = os.Create("./data/test.log")
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
