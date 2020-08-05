package linkedlist

import "fmt"

/*
	LRU算法
	双向链表 + hashmap
	2020.08.05 shopee命中一次
*/

/*
	为什么是双向链表？
	现在有四个元素  A --> D --> C --> E
	最近使用算法，需要维护每个元素根据最近使用的次序
	假如此时C被使用了，需要把C放到队列头部。

	数组不行。因为C放头部，后面的元素需要向后挪动
	单链表不行，因为C放头部，D需要指向E.我们需要从头遍历到D

	只有双向链表可以！
*/！
*/








type DataNode struct {
	Key  int // 需要额外保存key，因为从nodelist里无法推出一个node在map的key是啥
	Val  int
	Pre  *DataNode
	Next *DataNode
}

type LRUCache struct {
	data map[int]*DataNode
	len  int
	cap  int
	head *DataNode // 新数据要插到头部
	tail *DataNode // 淘汰数据从尾部删除 -- 最旧未使用
}

func Constructor(capacity int) LRUCache {
	lcache := new(LRUCache)
	lcache.head = new(DataNode)
	lcache.tail = new(DataNode)
	lcache.head.Next = lcache.tail
	lcache.tail.Pre = lcache.head

	lcache.data = make(map[int]*DataNode, capacity)
	lcache.cap = capacity

	return *lcache
}

func (this *LRUCache) Get(key int) int {
	fmt.Printf("start Get Key %d\n", key)
	v, ok := this.data[key]
	if !ok {
		return -1
	}
	// 需要两个工具函数
	this.remove(v)        // 先从链表中移出来
	this.addNodeToHead(v) // 放到头部
	return v.Val
}

func (this *LRUCache) Put(key int, value int) {
	d, ok := this.data[key]
	if ok {
		d.Val = value
		this.remove(d)        // 先移出来
		this.addNodeToHead(d) // 放到头部
		return
	}
	fmt.Printf("start put %d %d\n", key, value)

	// add new key
	if this.len >= this.cap { // del tail key
		last := this.tail.Pre
		this.remove(last)
		delete(this.data, last.Key) // map中删除
		this.len--
	}

	n := &DataNode{}
	n.Key = key
	n.Val = value
	this.addNodeToHead(n)

	this.data[key] = n
	this.len++
	// this.Print()
}

// 把节点放到头部
func (this *LRUCache) addNodeToHead(in *DataNode) {
	this.head.Next.Pre = in
	in.Next = this.head.Next
	in.Pre = this.head
	this.head.Next = in
}

func (this *LRUCache) remove(in *DataNode) {
	in.Next.Pre = in.Pre // 把结点脱离出来
	in.Pre.Next = in.Next
}

func (this *LRUCache) Print() {
	cur := this.head
	fmt.Println("Do Print:")
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}


