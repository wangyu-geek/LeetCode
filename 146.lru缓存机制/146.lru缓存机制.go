package _46_lru缓存机制

type LRUCache struct {
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity}
}

/**
 *
 */
func (this *LRUCache) Get(key int) int {

}

func (this *LRUCache) Put(key int, value int) {

}

type Node struct {
	key   int
	value int
	pre   Node
	next  Node
}

func ConstructorNode(key, value int) Node {
	return Node{
		key:   key,
		value: value,
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
