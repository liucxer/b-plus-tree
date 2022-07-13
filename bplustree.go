package b_plus_tree

const (
	// 5 node caches are needed at least for self, left and right sibling, sibling of sibling, parent and node seeking
	MIN_CACHE_NUM  int64 = 5
	INVALID_OFFSET       = 0xdeadbeef

	BPLUS_TREE_LEAF     int64 = 0
	BPLUS_TREE_NON_LEAF int64 = 1

	ADDR_STR_WIDTH int64 = 16
)

type key_t int

type list_head struct {
	prev *list_head
	next *list_head
}

func list_init(link *list_head) {
	link.prev = link
	link.next = link
}

func __list_add(link *list_head, prev *list_head, next *list_head) {
	prev.next = link

	link.prev = prev
	link.next = next

	next.prev = link
}

func __list_del(prev *list_head, next *list_head) {
	prev.next = next
	next.prev = prev
}

func list_add(link *list_head, prev *list_head) {
	__list_add(link, prev, prev.next)
}

func list_add_tail(link *list_head, head *list_head) {
	__list_add(link, head.prev, head)
}

func list_del(link *list_head) {
	__list_del(link.prev, link.next)
	list_init(link)
}

func list_empty(head *list_head) bool {
	return head.next == head
}

type bplus_node struct {
	self      int64
	parent    int64
	prev      int64
	next      int64
	node_type int64

	/* If leaf node, it specifies  count of entries,
	 * if non-leaf node, it specifies count of children(branches) */
	children int64
}

type free_block struct {
	link   list_head
	offset int64
}

type bplus_tree struct {
	caches      string
	used        [MIN_CACHE_NUM]int
	filename    string
	fd          int64
	level       int64
	root        int64
	file_size   int64
	free_blocks list_head
}

var (
	_block_size  int64
	_max_entries int64
	_max_order   int64
)

func is_leaf(node *bplus_node) bool {
	return node.node_type == BPLUS_TREE_LEAF
}

/*

#define offset_ptr(node) ((char *) (node) + sizeof(*node))
#define key(node) ((key_t *)offset_ptr(node))
#define data(node) ((long *)(offset_ptr(node) + _max_entries * sizeof(key_t)))
#define sub(node) ((off_t *)(offset_ptr(node) + (_max_order - 1) * sizeof(key_t)))
*/

func key_binary_search(node *bplus_node, target int64) {
	/*
		key_t *arr = key(node);
		int len = is_leaf(node) ? node->children : node->children - 1;
		int low = -1;
		int high = len;

		while (low + 1 < high) {
		int mid = low + (high - low) / 2;
		if (target > arr[mid]) {
		low = mid;
		} else {
		high = mid;
		}
		}

		if (high >= len || arr[high] != target) {
		return -high - 1;
		} else {
		return high;
		}
	*/
}
