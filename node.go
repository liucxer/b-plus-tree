package b_plus_tree

// B+树阶数
const (
	BTreeOrder = 5 // 一颗5阶B树的插入过程，5阶B数的结点最少2个key，最多4个key
)

type NodeType string

// 数据恢复状态
const (
	UnknownNodeType NodeType = "unknown"
	RootNodeType             = "root"     // 根节点
	NonLeafNodeType          = "non_leaf" // 非叶子节点
	LeafNodeType             = "leaf"     // 叶子节点

)

// 叶子节点
type Node struct {
	NodeType
	Key          [BTreeOrder - 1]int64
	RightBrother *Node
	Children     [BTreeOrder]*Node
	Data         string
}

// 当前节点是否满了
func (node *Node) NodeIsFull() bool {
	for _, v := range node.Key {
		if v == 0 {
			return false
		}
	}

	return true
}

// 当前节点的key的长度
func (node *Node) NodeLen() int64 {
	var length int64
	for i := 0; i < len(node.Key); i++ {
		if node.Key[i] != 0 {
			length++
		}
	}

	return length
}

// 找到应该插入的offset
func (node *Node) FindOffset() int64 {
	for i := int64(0); i < node.NodeLen(); i++ {

	}

	return 0
}

func (node *Node) NodeInsert(key int64, data string) {
	if node.NodeIsFull() {

	} else {

	}
}
