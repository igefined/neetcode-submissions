type MyLinkedList struct {
	head, tail *ListNode
	length     int
}

type ListNode struct {
	val  int
	next *ListNode
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
	this.head = &ListNode{val: val, next: this.head}
	this.length++
}

func (this *MyLinkedList) AddAtTail(val int) {
	next := this.head
	for i := 0; next != nil; i++ {
		if next.next == nil {
			next.next = &ListNode{val: val}
			this.length++
			break
		}
		next = next.next
	}

	//newNode := &ListNode{val: val}
	//prevTail := this.tail
	//this.tail = newNode
	//if prevTail != nil {
	//	prevTail.next = newNode
	//}
	//this.length++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
	}

	if index == this.length {
		this.AddAtTail(val)
		return
	}

	next := this.head
	var prev *ListNode
	for i := 0; next != nil; i++ {
		if i == index {
			prev.next = &ListNode{val: val, next: next}
			this.length++
		}
		prev = next
		next = next.next
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	next := this.head
	var prev *ListNode

	for i := 0; next != nil; i++ {
		if i == index {
			prev.next = next.next
			this.length = this.length - 1
			break
		}
		prev = next
		next = next.next
	}

}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */