package main

import "fmt"

func sum(nums ...int) int {
	sum := 0
	fmt.Printf("nums 타입: %T\n", nums)

	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {

}
