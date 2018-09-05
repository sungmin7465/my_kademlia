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

type KBucket struct {
	*list.List
}

type DHT struct {
	hostid   string
	kbuckets [KBucketSize]KBucket
}

type Node struct {
	id  string
	dht *DHT
}

func (dht DHT) Update(peer Node) {
	//거리에 맞는 bucket에 넣고 update
	index := GetBucketNumber(dht.hostid, peer.id)
	kb := dht.kbuckets[index]
	kb.Update(dht.hostid, peer)
}

func (kb KBucket) Update(hostid string, peer Node) {
	var found_e *list.Element

	for e := kb.Front(); e != nil; e = e.Next() {
		if hostid == e.Value.(Node).id {
			found_e = e
		}
	}

	if found_e != nil {
		// If entry is already in KBucket, move it to back of list
		kb.MoveToBack(found_e)
	} else if kb.Len() >= KBucketSize { // is Full
		// Ping node, and remove if unresponsive
		// TODO(@cfromknecht) Build internal ping

		//kb.Remove(bucket.Front())
		//kb.PushBack(foundPtr)
	} else {
		// KBucket is not full, simply add contact
		kb.PushBack(peer)
	}

}

func NewNodeRandomID() *Node {

	new_id := NewRandomNodeID()

	ret := &Node{
		id:  new_id,
		dht: NewDHT(new_id),
	}

	return ret
}

func NewNode(id string) *Node {

	new_id := id

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
		new_kbuckets[i] = KBucket{list.New()}
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

func Xor(node_id string, other_id string) []byte {
	decoded_node_id, _ := hex.DecodeString(node_id)
	decoded_other_id, _ := hex.DecodeString(other_id)
	var result []byte

	for i, element := range decoded_node_id {
		result = append(result, element^decoded_other_id[i])
	}

	//ret := hex.EncodeToString(result)

	return result

}

func GetBucketNumber(node_id string, other_id string) int {
	xored_id := Xor(node_id, other_id)
	offset := 0

	for index := 0; index < len(xored_id); index++ {
		offset = (len(xored_id) - 1 - index) * 8
		if xored_id[index] > 0 {
			for xored_id[index] > 0 {
				xored_id[index] >>= 1
				offset += 1

			}
			break
		}

	}

	return offset
}
