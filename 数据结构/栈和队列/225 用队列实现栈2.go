package main

import "fmt"

//LeetCode 225. 用队列实现栈

//使用队列实现栈的下列操作：
//push(x) -- 元素 x 入栈
//pop() -- 移除栈顶元素
//top() -- 获取栈顶元素
//empty() -- 返回栈是否为空

//注意:
// 你只能使用队列的基本操作-- 也就是 push to back, peek/pop from front, size, 和 is empty 这些操作是合法的。
// 你所使用的语言也许不支持队列。 你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
// 你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）。

// 方法2：一队列来实现, 这个更形象
//一个队列在模拟栈弹出元素的时候只要将队列头部的元素（除了最后一个元素外） 重新添加到队列尾部，
// 此时在去弹出元素就是栈的顺序了。

type MyStack struct {
	queue []int //创建一个队列
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{ //初始化
		queue: make([]int, 0),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	//添加元素
	this.queue = append(this.queue, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	n := len(this.queue) - 1 //判断长度
	for n != 0 {             //除了最后一个，其余的都重新添加到队列里
		val := this.queue[0]
		this.queue = this.queue[1:]
		this.queue = append(this.queue, val)
		n--
	}
	//弹出元素
	val := this.queue[0]
	this.queue = this.queue[1:]
	return val

}

/** Get the top element. */
func (this *MyStack) Top() int {
	//利用Pop函数，弹出来的元素重新添加
	val := this.Pop()
	this.queue = append(this.queue, val)
	return val
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

func main() {
	stack := Constructor()
	stack.Push(1)
	fmt.Println(stack.Top())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Empty())

}
