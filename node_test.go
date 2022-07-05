package b_plus_tree_test

import (
	b_plus_tree "bplustree"
	"github.com/davecgh/go-spew/spew"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func GetUniqueNum(count int64) []int64 {
	rand.Seed(time.Now().UnixNano())

	intMap := map[int64]int64{}
	for i := int64(0); i < count*2; i++ {
		randInt := int64(rand.Int63n(10000))
		intMap[randInt] = randInt
	}

	var res []int64
	for _, v := range intMap {
		res = append(res, v)
	}

	return res[0:count]
}

func TestNode_NodeInsertForLeafNode(t *testing.T) {
	node := b_plus_tree.Node{IsLeafNode: true}

	intSlice := GetUniqueNum(4)
	for _, item := range intSlice {
		node.NodeInsertForLeafNode(item, strconv.Itoa(int(item)))
	}
	spew.Dump(node)
}

func TestNode_NodeInsertForRootNode(t *testing.T) {
	node := &b_plus_tree.Node{IsLeafNode: true, IsRootNode: true}

	intSlice := GetUniqueNum(5)
	for _, item := range intSlice {
		node.NodeInsertForRootNode(item, strconv.Itoa(int(item)))
	}
	spew.Dump(node)
}
