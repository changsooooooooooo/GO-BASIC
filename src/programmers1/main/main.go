package main

import (
	"fmt"
	"programmers1/kthNumber"
)

func main(){
	arr:=kthNumber.Solution([7]int{1, 5, 2, 6, 3, 7, 4}, [3][3]int{{2, 5, 3}, {4, 4, 1}, {1, 7, 3}})
	fmt.Print(arr)
}
