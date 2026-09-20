type MyLinkedList struct {
	head, tail *ListNode
	length     int
}

type ListNode struct {
	val  int
	next *ListNode
	prev *ListNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index > this.length {
		return -1
	}

	next := this.head
	for i := 0; next != nil; i++ {
		if i == index {
			return next.val
		}
		next = next.next
	}

	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	newnode := &ListNode{val: val, next: this.head}
	if this.head != nil {
		this.head.prev = newnode
	} else {
		this.tail = newnode
	}
	this.head = newnode
	this.length++
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &ListNode{val: val, prev: this.tail}
	if this.tail != nil {
		this.tail.next = newNode
	} else {
		this.head = newNode
	}
	this.tail = newNode
	this.length++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}

	if index == this.length {
		this.AddAtTail(val)
		return
	}

	next := this.head
	var prev *ListNode
	for i := 0; i <= this.length; i++ {
		if i == index {
			newNode := &ListNode{val: val, next: next, prev: prev}
			prev.next = newNode
			if next != nil {
				next.prev = newNode
			}
			this.length++
			break
		}
		prev = next
		next = next.next
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		this.head = this.head.next
		this.length--
		if this.length == 0 {
			this.tail = nil
		}
		return
	}

	next := this.head
	var prev *ListNode

	for i := 0; i < this.length; i++ {
		if i == index {
			prev.next = next.next
			if next.next != nil {
				next.next.prev = prev
			}
			this.length--
			if index == this.length {
				this.tail = prev
			}
			break
		}
		prev = next
		next = next.next
	}
}