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


// use this to block the main process before exit
// thus make the raft server server indefinitely
var serverchan chan int

// start the server to receive other client(server)'s RPC
func StartServer(index int) {
	fmt.Println("Raft server start.")
	s := server.NewServer()
	s.Register(new(raft.Raft), "")
	// s.Serve("tcp", utility.PeerAddrs[index])
	s.Serve("tcp", raft.PeerAddrs[index])
	serverchan <- 1
}

func main () {
	// use sys args for test
	// TODO: should be changed to the server's real index in the addr[]
	index, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Need an int argument to indicate the index of the server.")
		return
	}
	go StartServer(index)

	// sleep explicitly to wait for other servers to start
	time.Sleep(10 * time.Second)

	// start a Raft server
	cfg := utility.Make_config(index)
	time.Sleep(3 * time.Second)
	cfg.ShowLeader()
	// cfg.One(os.Args[1])
	cfg.ShowLeader()
	for {
		// cfg.ShowLeader()
		cfg.GetLeader()
		time.Sleep(2 * time.Second)
	}


	<- serverchan

}
