package leetcode

type Node struct {
	prev  *Node
	next  *Node
	key   int
	value int
}

type LRUCache struct {
	keyToNode map[int]*Node
	cap       int
	head      *Node
	tail      *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{key: -1, value: -1}
	tail := &Node{key: -1, value: -1}
	c := LRUCache{
		cap:       capacity,
		keyToNode: make(map[int]*Node, capacity),
		head:      head,
		tail:      tail,
	}
	c.join(head, tail)
	return c
}

func (this *LRUCache) moveToHead(node *Node) {
	this.join(node, this.head.next)
	this.join(this.head, node)
}

func (this *LRUCache) join(node1 *Node, node2 *Node) {
	node1.next = node2
	node2.prev = node1
}

func (this *LRUCache) remove(node *Node) {
	this.join(node.prev, node.next)
}

func (this *LRUCache) Get(key int) int {
	node, isOk := this.keyToNode[key]
	if isOk {
		this.remove(node)
		this.moveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node, exist := this.keyToNode[key]
	if len(this.keyToNode) >= this.cap && !exist {
		delete(this.keyToNode, this.tail.prev.key)
		this.remove(this.tail.prev)
	}
	if exist {
		this.remove(node)
	}
	node = &Node{key: key, value: value}
	this.moveToHead(node)
	this.keyToNode[key] = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
