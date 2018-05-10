package utility

import (
	"../raft"
	"sync"
)

// Global(static) server peers addr array
//var PeerAddrs = [...]string{"localhost:8971",
//                        "localhost:8972",
//                        "localhost:8973", }
//
//var Peer_n = 3

// save each Raft server instance
//type ClientEnd struct {
//	// endname = PeerAddrs[endindex]
//	endIndex int // corresponding to the peers array
//	addr string
//}

type Config struct {
	mu sync.Mutex
	saved *raft.Persister
	endIndex int
	rf *raft.Raft
	applyCh chan raft.ApplyMsg
}

//type ApplyMsg struct {
//	Index       int
//	Command     interface{}
//	UseSnapshot bool   // ignore for lab2; only used in lab3
//	Snapshot    []byte // ignore for lab2; only used in lab3
//}

// return the ClientEnd peers as 'raft.Make' parameter
func getClientPeers() []*raft.ClientEnd{
	cep := make([]*raft.ClientEnd, raft.Peer_n)
	for i := 0; i < raft.Peer_n; i++ {
		cep[i] = &raft.ClientEnd{}
		cep[i].EndIndex = i
		cep[i].Addr = raft.PeerAddrs[i]
	}
	return cep
}

func Make_config(index int) *Config {
	cfg := &Config{}
	cfg.saved = raft.MakePersister()
	cfg.endIndex = index
	cfg.applyCh = make(chan raft.ApplyMsg)
    // func Make(peers []*utility.ClientEnd, me int,
    //     persister *Persister, applyCh chan ApplyMsg) *Raft {}
	cfg.rf = raft.Make(getClientPeers(), index, cfg.saved, cfg.applyCh)

	return cfg
}


func (cfg *Config) One(cmd interface{}) {
	// TODO: wrapper for rf.Start(cmd)
	cfg.rf.Start(cmd)
}

func (cfg *Config) ShowLog() {
	// TODO: show logs that committed in this raft instance
}

// rpc call example from 'labrpc'
//	ok := rf.peers[server].Call("AppendEntries", args, reply)

// Call inside Raft instance, to revoke an RPC to Raft instance on other server
// send an RPC, wait for the reply,
// and return value indicates success; false means the server couldn't be contacted.

//func (e *ClientEnd) Call(svcMeth string, args interface{}, reply interface{}) bool {
//	// call self
//	// if e.endIndex == cfg.endIndex {
//		//return false
//	// }
//	d := client.NewPeer2PeerDiscovery("tcp@" + PeerAddrs[e.endIndex], "")
//	xclient := client.NewXClient("Raft", client.Failtry, client.RandomSelect, d, client.DefaultOption)
//	err := xclient.Call(context.Background(), svcMeth, args, reply)
//	defer xclient.Close()
//	if err != nil {
//		return false
//	}
//	// err := xclient.Call(context.Background(), "Dull", &args, &reply)
//
//	return true
//}
