package kthNumber

import (
	"fmt"
	"sort"
)

func Solution(array [7]int, commands [3][3]int) []int {

	var answer []int
	for i:=0; i< len(commands); i++{
		tempArray:=array[commands[i][0]-1:commands[i][1]]
		fmt.Println(tempArray)
		sort.Ints(tempArray)
		answer = append(answer, tempArray[commands[i][2]-1])
	}

	return answer
}
