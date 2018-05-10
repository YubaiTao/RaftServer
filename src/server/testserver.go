package main

import (
	"../raft"
	"fmt"
)

func main () {
	client := &raft.ClientEnd{}
	client.EndIndex = 0
	client.Addr = raft.PeerAddrs[0]
	str := "test string"
	var reply_ string
	args_ := str

	var args raft.AppendEntriesArgs
	args.Term = 0
	args.LeaderId = 0
	args.PrevLogIndex = 0
	args.PrevLogTerm = 0
	args.Entry = nil
	args.LeaderCommit = 0

	// var reply raft.AppendEntriesReply
	reply := raft.AppendEntriesReply{}
	reply.NextIndex = 0
	reply.Success = false
	reply.Term = 0



	client.Call("Dull", args_, &reply_)
	fmt.Println("reply: " , reply_)

	fmt.Println("--------------")
	client.Call("AppendEntries", args, &reply)
	fmt.Println("NextIndex: ", reply.NextIndex)

}
