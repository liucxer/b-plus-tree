package b_plus_tree

import "fmt"

// B+树阶数
const (
	BTreeOrder     = int64(5) // 一颗5阶B树的插入过程，5阶B数的结点最少2个key，最多4个key
	MaxKeyLen      = BTreeOrder - 1
	MaxChildrenLen = BTreeOrder
)

// 叶子节点
type Node struct {
	IsRootNode    bool // 根节点
	IsNonLeafNode bool // 非叶子节点
	IsLeafNode    bool // 叶子节点

	KenLen       int64
	Key          [MaxKeyLen]int64
	RightBrother *Node
	Children     [MaxChildrenLen]*Node
	Data         [MaxKeyLen]string
}

// 当前节点是否满了
func (node *Node) NodeIsFull() bool {
	if node.KenLen == MaxKeyLen {
		return true
	}
	return false
}

// 根节点插入数据(未满场景)
func (node *Node) NodeInsertForRootNode(key int64, data string) {
	if node.NodeIsFull() {
		var (
			tmpKeys     [MaxKeyLen + 1]int64
			tmpChildren [MaxChildrenLen + 1]*Node
			tmpDatas    [MaxKeyLen + 1]string
		)
		for i := int64(0); i < MaxKeyLen; i++ {
			tmpKeys[i] = node.Key[i]
			tmpDatas[i] = node.Data[i]
		}
		for i := int64(0); i > MaxChildrenLen; i++ {
			tmpChildren[i] = node.Children[i]
		}

		i := int64(0)
		for ; i < MaxKeyLen; i++ {
			if tmpKeys[i] < key {
				continue
			}
			for j := MaxKeyLen; j > i; j-- {
				tmpKeys[j] = tmpKeys[j-1]
				tmpDatas[j] = tmpDatas[j-1]
			}
			break
		}

		tmpKeys[i] = key
		tmpDatas[i] = data
		fmt.Println(tmpKeys)
		fmt.Println(tmpDatas)

		rightChildrenNode := Node{
			IsNonLeafNode: true,
			IsLeafNode:    true,
		}
		for i := (MaxKeyLen + 1) / 2; i < MaxKeyLen+1; i++ {
			rightChildrenNode.Key[i-((MaxKeyLen+1)/2)] = tmpKeys[i]
			rightChildrenNode.Data[i-((MaxKeyLen+1)/2)] = tmpDatas[i]
		}

		leftChildrenNode := Node{
			IsNonLeafNode: true,
			IsLeafNode:    true,
		}
		for i := int64(0); i < (MaxKeyLen+1)/2; i++ {
			leftChildrenNode.Key[i] = tmpKeys[i]
			leftChildrenNode.Data[i] = tmpDatas[i]
		}
		leftChildrenNode.RightBrother = &rightChildrenNode

		fatherNode := Node{
			IsNonLeafNode: true,
			IsLeafNode:    true,
		}
		fatherNode.Key[0] = tmpKeys[(MaxKeyLen+1)/2]
		fatherNode.Children[0] = &leftChildrenNode
		fatherNode.Children[1] = &rightChildrenNode
		fatherNode.Data[0] = tmpDatas[(MaxKeyLen+1)/2]

		node.KenLen = 1
		node.Key = fatherNode.Key
		node.Children = fatherNode.Children
		node.Data = fatherNode.Data
	} else {
		if node.IsLeafNode {
			node.NodeInsertForLeafNode(key, data)
		}
	}
}

// 叶子节点插入数据(未满场景)
func (node *Node) NodeInsertForLeafNode(key int64, data string) {
	if node.NodeIsFull() {
		panic("LeafNode is full")
	}

	i := int64(0)
	for ; i < node.KenLen; i++ {
		if node.Key[i] < key {
			continue
		}
		for j := node.KenLen; j > i; j-- {
			node.Key[j] = node.Key[j-1]
			node.Data[j] = node.Data[j-1]
		}
		break
	}

	node.Key[i] = key
	node.Data[i] = data

	node.KenLen++
}
