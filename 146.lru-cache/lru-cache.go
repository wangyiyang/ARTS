package main

type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

type LRUCache struct {
	limit int
	hash  map[int]*Node
	Head  *Node
	Tail  *Node
}

func Constructor(capacity int) LRUCache {
	h := &Node{-1, -1, nil, nil}
	t := &Node{-1, -1, nil, nil}
	h.Next = t
	t.Prev = h
	hash := make(map[int]*Node, capacity)
	cache := LRUCache{hash: hash, limit: capacity, Head: h, Tail: t}
	return cache
}

func (this *LRUCache) insert(node *Node) {
	t := this.Tail
	node.Prev = t.Prev
	t.Prev.Next = node
	node.Next = t
	t.Prev = node
}

func (this *LRUCache) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.hash[key]; ok {
		this.remove(v)
		this.insert(v)
		return v.Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.hash[key]; ok {
		this.remove(v)
		this.insert(v)
		v.Val = value
	} else {
		if len(this.hash) >= this.limit {
			h := this.Head.Next
			this.remove(h)
			delete(this.hash, h.Key)
		}
		node := &Node{key, value, nil, nil}
		this.hash[key] = node
		this.insert(node)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {

}
