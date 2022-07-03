package b_plus_tree

type BPlusTree Node

func NewBPlusTree() *BPlusTree {
	var bPlusTree BPlusTree
	bPlusTree.NodeType = RootNodeType
	return &bPlusTree
}

func (tree *BPlusTree) Insert(key int64, data string) {

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
