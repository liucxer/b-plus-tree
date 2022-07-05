package b_plus_tree

type BPlusTree struct {
	*Node
}

func NewBPlusTree() *BPlusTree {
	var bPlusTree BPlusTree

	var node Node
	node.IsRootNode = true
	node.IsNonLeafNode = true
	node.IsLeafNode = true
	bPlusTree.Node = &node
	return &bPlusTree
}

func (tree *BPlusTree) Insert(key int64, data string) {
	if tree.IsLeafNode && !tree.NodeIsFull() {
		tree.Node.NodeInsertForLeafNode(key, data)
	}
}

func (tree *BPlusTree) Delete(key int64) {

}

func (tree *BPlusTree) List(key int64) []*Node {
	var (
		res []*Node
	)

	return res
}

func (tree *BPlusTree) Range(start, end int64) []*Node {
	var (
		res []*Node
	)

	return res
}
