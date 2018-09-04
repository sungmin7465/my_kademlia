package main

import (
	"fmt"
	"my_kademlia/kademlia"
)

func main() {
	node := kademlia.NewNode()
	fmt.Println(node)
}
