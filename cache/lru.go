package cache

type node struct {
	prev  *node
	next  *node
	key   interface{}
	value interface{}
}

type LRU struct {
	head *node
	tail *node
	len  int64
	cap  int64
	kv   map[interface{}]*node
}

func NewLRUCache(cap int64) *LRU {
	if cap < 0 {
		panic("invalid cap")
	}
	head := &node{}
	tail := &node{}
	head.next = tail
	tail.prev = head
	return &LRU{
		head: head,
		tail: tail,
		cap:  cap,
		kv:   make(map[interface{}]*node),
	}
}

func (l *LRU) Get(key interface{}, defaultValue interface{}) interface{} {
	if node := l.getNode(key); node != nil {
		l.moveNode(node)
		return node.value
	}
	return defaultValue
}

func (l *LRU) Set(key interface{}, value interface{}) {
	if l.Contains(key) {
		l.removeNode(key)
	}
	if l.len >= l.cap {
		l.removeLastNode()
	}
	l.addNode(key, value)
}

func (l *LRU) Contains(key interface{}) bool {
	return l.getNode(key) == nil
}

func (l *LRU) Remove(key interface{}) {
	if !l.Contains(key) {
		return
	}

	l.removeNode(key)
}

func (l *LRU) GetLen() int64 {
	return l.len
}

func (l *LRU) getNode(key interface{}) *node {
	if node, found := l.kv[key]; found {
		return node
	}
	return nil
}

func (l *LRU) removeNode(key interface{}) {
	node := l.getNode(key)
	if node == nil {
		return
	}

	prevNode := node.prev
	nextNode := node.next

	prevNode.next = nextNode
	nextNode.prev = prevNode

	delete(l.kv, key)
	l.len--
}

func (l *LRU) addNode(key interface{}, value interface{}) {
	newNode := &node{
		key:   key,
		value: value,
	}

	nextNode := l.head.next
	l.head.next = newNode
	newNode.prev = l.head
	newNode.next = nextNode

	l.kv[key] = newNode
	l.len++
}

func (l *LRU) moveNode(node *node) {
	if node == nil || node.prev == l.head {
		return
	}

	pNode := node.prev
	nNode := node.next
	pNode.next = nNode
	nNode.prev = pNode

	fNode := l.head.next
	l.head.next = node
	node.prev = l.head
	node.next = fNode
	fNode.prev = node
}

func (l *LRU) removeLastNode() {
	node := l.tail.prev
	if node == nil || node == l.head {
		return
	}
	l.removeNode(node.key)
}
