type MinStack struct {
    Items[]int
    Mins[]int
}

func Constructor() MinStack {
    var st MinStack
    return st
}

func (this *MinStack) Push(val int) {
    this.Items = append(this.Items,val)
    n := len(this.Mins)
    if n == 0 || val < this.Mins[n-1] {
        this.Mins = append(this.Mins,val)
    } else {
        this.Mins = append(this.Mins, this.Mins[n-1])
    }
}

func (this *MinStack) Pop() {
    n := len(this.Items) - 1
    this.Items = this.Items[:n]
    this.Mins = this.Mins[:n]
}

func (this *MinStack) Top() int {
    return this.Items[len(this.Items)-1]
}

func (this *MinStack) GetMin() int {
    return this.Mins[len(this.Mins)-1]
}
