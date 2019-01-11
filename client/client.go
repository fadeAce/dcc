package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/monnand/dhkx"
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
	fmt.Println(*conf)
}

func main() {
	args := os.Args
	if len(args) == 3 {
		node := args[1]
		cmd := args[2]
		do := node == "all"
		for _, val := range conf.Nodes {
			if do {
				// do cmd sending

				continue
			}
			if node == val.Name {
				// cmd sending

			}
		}
		fmt.Println(node, "exec", cmd)
	}

}

func send(cmd string) {

}

func makePost() {

}

func dhkxfound() {
	g, _ := dhkx.GetGroup(0)
	priv, _ := g.GeneratePrivateKey(nil)
	pub := priv.Bytes()

	// Recover Bob's public key
	bobPubKey := NewPublicKey(b)

	// Compute the key
	k, _ := g.ComputeKey(bobPubKey, priv)

	// Get the key in the form of []byte
	key := k.Bytes()
}
