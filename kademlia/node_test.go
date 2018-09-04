package kademlia

import (
	"fmt"
	"testing"
)

func TestPrintId(t *testing.T) {
	id := NewRandomNodeID()

	fmt.Println(id)
}
