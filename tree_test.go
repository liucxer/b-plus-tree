package b_plus_tree_test

import (
	b_plus_tree "bplustree"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestNode_Insert(t *testing.T) {
	bPlusTree := b_plus_tree.NewBPlusTree()

	bPlusTree.Insert(10, "10")
	spew.Dump(bPlusTree)
}
