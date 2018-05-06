package main

import (
	"../utility"
	"../raft"
	"os"
	"strconv"
	"fmt"
	"github.com/smallnest/rpcx/server"
	"time"
)


// start the server to receive other client(server)'s RPC
func StartServer(index int) {
	s := server.NewServer()
	s.Register(new(raft.Raft), "")
	s.Serve("tcp", utility.PeerAddrs[index])
	fmt.Println("Server start.")
}

func main () {
	//rf := raft.Make()

	index, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Need an int argument to indicate the index of the server.")
		return
	}
	go StartServer(index)

	// sleep explicitly to wait for other servers to start
	time.Sleep(5 * time.Second)

	cfg := utility.Make_config(0)
	cfg.One("test1")

	// func Make(peers []*utility.ClientEnd, me int,
		// persister *Persister, applyCh chan ApplyMsg) *Raft {


}
