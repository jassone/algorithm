package main

import "fmt"

//LeetCode 232.用栈实现队列

//使用栈实现队列的下列操作：

//push(x) -- 将一个元素放入队列的尾部。
//pop() -- 从队列首部移除元素。
//peek() -- 返回队列首部的元素。
//empty() -- 返回队列是否为空。

// 卡尔版
// 使用栈来模式队列的行为，如果仅仅用一个栈，是一定不行的，所以需要两个栈一个输入栈，一个输出栈，
// 这里要注意输入栈和输出栈的关系。

type MyQueue struct {
	stack []int
	back  []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stack: make([]int, 0),
		back:  make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	// 入栈，所以需要把栈b数据再挪到栈a里面，然后再添加当前元素
	for len(this.back) != 0 {
		val := this.back[len(this.back)-1]
		this.back = this.back[:len(this.back)-1]
		this.stack = append(this.stack, val)
	}
	this.stack = append(this.stack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	// 出栈，所以要把栈a数据反过来放到另一个栈b里面，然后输出b的第一个元素即可
	for len(this.stack) != 0 {
		val := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		this.back = append(this.back, val)
		//fmt.Println("bb")
	}
	if len(this.back) == 0 {
		return 0
	}
	val := this.back[len(this.back)-1]
	this.back = this.back[:len(this.back)-1]
	return val
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	// 先拿出来
	val := this.Pop()
	if val == 0 {
		return 0
	}
	// 再放进去
	this.back = append(this.back, val)
	return val
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0 && len(this.back) == 0
}

func main() {
	queue := Constructor()

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Pop()
	fmt.Println(queue)

	fmt.Println(queue.Empty())
	fmt.Println(queue.Peek())
}
