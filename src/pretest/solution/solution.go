package solution

func Solution(answer []int) []int{
	first:=[...]int{1, 2, 3, 4, 5}
	second:=[...]int{2, 1, 2, 3, 2, 4, 2, 5}
	third:=[...]int{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}

	correct:=make([]int, 3)

	for i, v := range answer{
		if first[i%len(first)] == v{
			correct[0] += 1
		}

		if second[i%len(second)] == v{
			correct[1] += 1
		}

		if third[i%len(third)] == v{
			correct[2] += 1
		}
	}

	max:=0

	for _, v:=range correct{
		if max<v{
			max = v
		}
	}

	candidate:=make([]int, 0, 3)

	for i, _ := range correct{
		if correct[i] == max{
			candidate=append(candidate, i+1)
		}
	}
	return candidate
}
