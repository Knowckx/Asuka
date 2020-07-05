package struct

// 设计一个栈，这个栈可以在常数时间内检索到目前的最小值

// 解决：起一个新栈，专门用来保存目前栈顶的数所对应的最小值


type MinStack struct{
	Datas []int
	Len int
	Mins []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	out := MinStack{}
	out.Datas = make([]int,0,8)
	out.Mins = make([]int,0,8)
	return out
}



func (this *MinStack) Push(x int)     {
	this.Datas = append(this.Datas,x)
	this.Len++

	if len(this.Mins) == 0     {
		this.Mins = append(this.Mins,x)
		return
	}

	min := this.Mins[len(this.Mins)-1]
	if x < min {
		min = x
	}
	this.Mins = append(this.Mins,min)
}


func (this *MinStack) Pop()  {
	if this.Len == 0 {
		return 
	}
	this.Datas = this.Datas[:this.Len-1]
	this.Len--

	this.Mins = this.Mins[:len(this.Mins)-1]
}


func (this *MinStack) Top() int {

	return this.Datas[this.Len-1]

}


func (this *MinStack) GetMin() int {
	return this.Mins[this.Len-1]
}

