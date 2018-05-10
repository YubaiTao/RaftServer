package main

import (
	"raft"
	"fmt"
)

func main () {
	client := &raft.ClientEnd{}
	client.EndIndex = 0
	client.Addr = raft.PeerAddrs[0]
	str := "test string"
	var reply string
	args := str

	client.Call("Dull", args, &reply)
	fmt.Println("reply: " , reply)
}
