package kademlia

import (
	"container/list"
	"crypto/rand"
	"encoding/hex"
)

const (
	IDLength    = 32 //임시로 32. spec은 160
	KBucketSize = 32
)

type KBucket *list.List

type DHT struct {
	hostid   string
	kbuckets [KBucketSize]KBucket
}

type Node struct {
	id  string
	dht *DHT
}

func NewNode() *Node {

	new_id := NewRandomNodeID()

	ret := &Node{
		id:  new_id,
		dht: NewDHT(new_id),
	}

	return ret
}

func NewDHT(id string) *DHT {

	ret := &DHT{
		hostid:   id,
		kbuckets: NewKBuckets(),
	}

	return ret
}

func NewKBuckets() [KBucketSize]KBucket {

	var new_kbuckets [KBucketSize]KBucket

	for i := 0; i < KBucketSize; i++ {
		new_kbuckets[i] = list.New()
	}

	return new_kbuckets
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func NewRandomNodeID() string {
	hex_str, _ := randomHex(IDLength / 8)
	return hex_str
}

/*
func (node string) Xor(other string) (ret string) {

	return
}
*/
