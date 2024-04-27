package main

import "fmt"

type Lists struct {
	key  int
	val  int
	pre  *Lists
	next *Lists
}

type LRUCache struct {
	limit   int
	count   int
	head    *Lists
	tail    *Lists
	element map[int]*Lists
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{}
	head := new(Lists)
	tail := new(Lists)
	lru.head = head
	lru.tail = tail
	head.next = tail
	tail.pre = head
	lru.limit = capacity
	lru.element = make(map[int]*Lists)
	return lru
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.element[key]; ok {
		this.MoveToFront(key)
		return this.element[key].val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.element[key]; ok {
		this.MoveToFront(key)
		this.element[key].val = value
		return
	}
	if this.count < this.limit {
		this.AddHead(key, value)
		return
	}
	this.RemoveTail()
	this.AddHead(key, value)
}

func (this *LRUCache) AddHead(key int, value int) {
	node := new(Lists)
	node.key = key
	node.val = value
	tmp := this.head.next
	this.head.next = node
	node.pre = this.head
	node.next = tmp
	tmp.pre = node
	this.count++
	this.element[key] = node
}

func (this *LRUCache) MoveToFront(key int) {
	node := this.element[key]
	node.pre.next = node.next
	node.next.pre = node.pre
	tmp := this.head.next
	this.head.next = node
	node.pre = this.head
	node.next = tmp
	tmp.pre = node
}

func (this *LRUCache) RemoveTail() {
	tmp := this.tail.pre.pre
	tmp.next = this.tail
	this.tail.pre = tmp
	delete(this.element, this.tail.pre.key)
	this.count--
}

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	obj.Get(2)
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))
}
