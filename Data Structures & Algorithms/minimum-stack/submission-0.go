type MinStack struct {
	stack 		[]int
	minIndexes 	[]int
}

func Constructor() MinStack {
	return MinStack{
		stack: 		make([]int, 0, 0),
		minIndexes: make([]int, 0, 0),
	}
}

func (this *MinStack) Push(val int) {
	lastMinIndex := 0
	if len(this.minIndexes) > 0 {
		lastMinIndex = this.minIndexes[len(this.minIndexes)-1]
	}

	if len(this.stack) == 0 || val < this.stack[lastMinIndex] {
		this.minIndexes = append(this.minIndexes, len(this.stack))
	}

	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	lastMinIndexes := this.minIndexes[len(this.minIndexes)-1]
	if lastMinIndexes == len(this.stack)-1 {
		this.minIndexes = this.minIndexes[:len(this.minIndexes)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack) - 1]
}

func (this *MinStack) GetMin() int {
	return this.stack[this.minIndexes[len(this.minIndexes)-1]]
}
