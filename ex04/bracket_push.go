package brackets

type Stack struct {
	x []byte
	y int
}

func Struct_adress() *Stack {
	return &Stack{}
}
func (point *Stack) Push(date byte) {
	point.x = append(point.x, date)
	point.y = point.y + 1
}
func (point *Stack) Pop() byte {
	res := point.x[(point.y - 1)]
	point.x = point.x[:point.y-1]
	point.y = point.y - 1
	return res
}
func (point *Stack) Read_lasts() byte {
	res := point.x[(point.y - 1)]
	return res
}
func (point *Stack) Location() int {
	return point.y
}
func Bracket(date string) (bool, error) {
	var stack *Stack = Struct_adress()
	if len(date) == 0 {
		return true, nil
	}
	for i := 0; i < len(date); i++ {
		if date[i] == '{' || date[i] == '[' || date[i] == '(' {
			stack.Push(date[i])
		} else if date[i] == '}' || date[i] == ']' || date[i] == ')' {
			if date[i] == '}' && stack.Location() > 0 {
				if stack.Read_lasts() == '{' {
					stack.Pop()
				} else {
					return false, nil
				}
			}
			if date[i] == ']' && stack.Location() > 0 {
				if stack.Read_lasts() == '[' {
					stack.Pop()
				} else {
					return false, nil
				}
			}
			if date[i] == ')' && stack.Location() > 0 {
				if stack.Read_lasts() == '(' {
					stack.Pop()
				} else {
					return false, nil
				}
			}
		}
	}
	if stack.Location() == 0 {
		return true, nil
	}
	return false, nil
}

/*func main () {
 fmt.Println(Check("w{jke[r]nf)ed}f"))
}*/
