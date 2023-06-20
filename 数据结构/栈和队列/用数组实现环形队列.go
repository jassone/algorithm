package main

import (
	"errors"
	"fmt"
	"os"
)

var maxSize = 3

type CircleQueue struct {
	maxSize int
	array   [3]int //数组
	head    int    // 指向队首 0
	tail    int    // 指向队尾 0
}

//判断队列是否已满
func (this *CircleQueue) IsFull() bool {
	fmt.Println("tail::",this.tail)
	return (this.tail+1)%this.maxSize == this.head
}

//判断队里是否为空
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

//进队列
func (this *CircleQueue) AddQueue(val int) (err error) {
	if this.IsFull() {
		return errors.New("队列已满")
	}
	//把值给尾部
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize
	return
}

//出队列
func (this *CircleQueue) GetQueue() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("队列已空")
	}
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return
}

//显示队列元素
func (this *CircleQueue) ListQueue() {
	fmt.Println("队列情况如下：")
	//计算出队列多少元素
	//比较关键的一步
	size := (this.tail + this.maxSize - this.head) % this.maxSize
	if size == 0 {
		fmt.Println("队列已空")
	}
	//定义一个辅助变量 指向head
	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, this.array[tempHead])
		// 关键
		tempHead = (tempHead + 1) % this.maxSize
	}
	fmt.Println()
}

//主函数
func main() {
	//初始化一个队列
	queue := &CircleQueue{
		maxSize: maxSize,
		head:    0,
		tail:    0,
	}
	var key string
	var val int
	for {
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从队列获取数据")
		fmt.Println("3. 输入show 表示显示队列")
		fmt.Println("4. 输入exit 表示显示队列")

		fmt.Print("请输入:")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {

				fmt.Println("加入队列ok")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数=", val)
			}
		case "show":
			queue.ListQueue()
		case "exit":
			os.Exit(0)
		}
	}
}