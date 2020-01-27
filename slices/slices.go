package main

import "fmt"

func main() {
	/*Slices*/
	s := make([]int, 10)
	s[0] = 10
	fmt.Println(s)
	/*Array*/
	s1 := [4]int{1, 2, 3, 4}
	s1 = append(s1, 5)
	fmt.Println(s1)
}
