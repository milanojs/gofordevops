package sum_test

import (
	"fmt"
)

func ExampleInts() {
	s := sum.Ints(1, 2, 3, 4, 5)
	fmt.Println("Sum of one to five is ", s)

}
