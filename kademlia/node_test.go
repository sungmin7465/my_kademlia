package kademlia

import (
	"fmt"
	"testing"
)

func TestPrintRandomId(t *testing.T) {
	id := NewRandomNodeID()
	fmt.Println(id)
}

/*
func TestXor(t *testing.T) {
	expectedXor := "a31d980e"
	id1_str := "f7a4d5cc"
	id2_str := "54b94dc2"

	if Xor(id1_str, id2_str) != expectedXor {
		t.Error(fmt.Sprintf("XOR of %s and %s should be %s", id1_str, id2_str, expectedXor))
	}
	fmt.Println(expectedXor)
}
*/
func TestXor(t *testing.T) {

	id1_str := "0002222c"
	id2_str := "01234dc2"
	fmt.Println(GetBucketNumber(id1_str, id2_str))

}
