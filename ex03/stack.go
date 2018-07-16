package stack

type Stack struct {
	x []int
	y int
}

func New() *Stack {
	return &Stack{}
}
func (point *Stack) Push(data int) {
	point.x = append(point.x, data)
	point.y = point.y + 1
}
func (point *Stack) Pop() int {
	res := point.x[(point.y - 1)]
	point.x = point.x[:point.y-1]
	point.y = point.y - 1
	return res
}
