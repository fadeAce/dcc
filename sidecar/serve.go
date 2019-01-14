package sidecar

import (
	"encoding/json"
	"fmt"
	"github.com/fadeAce/dcc/types"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// hello world, the web server
func handle(w http.ResponseWriter, req *http.Request) {
	// using DH exchanging to share secret
	if req.Method == "POST" {
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Printf("ioutil.ReadAll() error: %v\n", err)
			return
		}
		cmd := &types.Command{}
		err = json.Unmarshal(data, cmd)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("read resp.Body successfully:\n%v\n", cmd)
		// open dat file
		// do 1. append one cmd line there
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		// case using golang image
		if dir == "/go" {
			dir = "/usr/local/bin"
		}
		// case using ubuntu image
		if dir == "/" {
			dir = "/usr/local/bin"
		}
		err = appendToFile(dir+"/data/sock.dat", cmd.Cmd+"\n")
		if err != nil {
			_, _ = io.WriteString(w, err.Error())
			return
		}
		_, _ = io.WriteString(w, "done !")
	}
}

func Serve() {
	http.HandleFunc("/cmd", handle)
	err := http.ListenAndServe(":2510", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func appendToFile(fileName string, content string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("file create failed. err: " + err.Error())
	} else {
		n, _ := f.Seek(0, io.SeekEnd)
		_, err = f.WriteAt([]byte(content), n)
	}
	defer f.Close()
	return err
}
