package main

import (
	"fmt"
	"my_kademlia/kademlia"
)

func main() {
	//node := kademlia.NewNodeRandomID()
	node := kademlia.NewNode("54b94dc2")
	fmt.Println(node)
}
