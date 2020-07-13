package sum

//Ints func exported
func Ints(vs ...int) int {
	return ints(vs)
}

func ints(vs []int) int {
	if len(vs) == 0 {
		return 0
	}
	return ints(vs[1:]) + vs[0]
}
