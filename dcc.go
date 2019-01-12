package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fadeAce/dcc/sidecar"
	"github.com/fadeAce/dcc/types"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	Nodes []Nodes
}

type Nodes struct {
	Addr string
	Name string
}

var conf *Config

func init() {
	data, _ := ioutil.ReadFile("dcc.yml")
	conf = &Config{}
	err := yaml.Unmarshal(data, conf)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	args := os.Args
	if len(args) > 2 {
		node := args[1]
		cmdArr := args[2:]
		cmd := strings.Trim(fmt.Sprint(cmdArr), "[]")
		do := node == "all"
		sent := false
		for _, val := range conf.Nodes {
			if do {
				// do cmd sending
				send(val.Addr, cmd)
				if !sent {
					sent = true
				}
				continue
			}
			if node == val.Name {
				// cmd sending
				send(val.Addr, cmd)
				if !sent {
					sent = true
				}
			}
		}
		fmt.Println("remote", node, "exec:", cmd)
		if sent {
			// has been sent
		} else {
			fmt.Println("wrong match node")
		}
	} else if len(args) == 2 {
		do := args[1]
		if do == "serve" {
			// it serve certain container
			sidecar.Serve()
		}
		if do == "halt" {
			// it halt certain container
		}
		if do == "test" {
			// begin to test case

		}
	}
}

func send(addr, cmd string) {
	cmdReq := types.Command{
		Cmd: cmd,
	}
	homeURL := "http://" + addr + ":2510/cmd"

	buff, err := json.Marshal(cmdReq)
	req, err := http.NewRequest("POST", homeURL, bytes.NewBuffer(buff))
	if err != nil {
		fmt.Printf("http.NewRequest() error: %v\n", err)
		return
	}

	req.Header.Add("Content-Type", "application/json")

	c := &http.Client{}
	resp, err := c.Do(req)
	res, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(res))
	if err != nil {
		fmt.Printf("http.Do() error: %v\n", err)
		return
	}
	_ = resp.Body.Close()
}

// todo : encryption through peer to peer
func dhkxfound() {
	//g, _ := dhkx.GetGroup(0)
	//priv, _ := g.GeneratePrivateKey(nil)
	//pub := priv.Bytes()
	//
	//// Recover Bob's public key
	//bobPubKey := NewPublicKey(b)
	//
	//// Compute the key
	//k, _ := g.ComputeKey(bobPubKey, priv)
	//
	//// Get the key in the form of []byte
	//key := k.Bytes()
}
