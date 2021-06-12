package kthNumber

import (
	"sort"
)

func Solution(array []int, commands [][]int) []int {

	var answer []int
	for i:=0; i< len(commands); i++{
		tempArray:=append([]int{}, array[commands[i][0]-1:commands[i][1]]...)
		sort.Ints(tempArray)
		answer = append(answer, tempArray[commands[i][2]-1])
	}

	return answer
}
