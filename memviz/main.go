package main

import (
	"bytes"
	"io/ioutil"
	"os/exec"

	"github.com/bradleyjkemp/memviz"
)

// A fullnode in the chain
type Address struct {
	ip   string
	port int
}

type Node struct {
	endPoint string
	peers    []*Node
}

type Graph struct {
	nodes map[Address]*Node
}

func main() {
	buf := &bytes.Buffer{}

	node0 := Node{endPoint: "0:0"}
	node1 := Node{endPoint: "1:1"}
	node2 := Node{endPoint: "2:2"}

	node0.peers = append(node0.peers, &node1)
	node0.peers = append(node0.peers, &node2)
	node1.peers = append(node1.peers, &node0)
	node1.peers = append(node1.peers, &node2)
	node2.peers = append(node2.peers, &node0)
	node2.peers = append(node2.peers, &node1)

	memviz.Map(buf, nil)

	// Write the parsed data to disk
	id := "test"
	fileName := "/tmp/g-" + id
	outputName := "/tmp/rendered-g-" + id + ".png"
	err := ioutil.WriteFile(fileName, buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("sfdp", "-Tpng", fileName, "-o", outputName)
	cmd.Run()

	opCmd := exec.Command("open", outputName)
	opCmd.Run()
}
