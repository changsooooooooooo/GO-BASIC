package main

import (
	sol "game1/alg/solution"
)

func main() {

	number := sol.MakeRandom()
	count := 1
	sel := sol.MakeNewInputValue()
	answer := sol.CheckCorrect(count, sel, number)

	for {
		if answer == count {
			break
		}
		count = answer
		sel = sol.MakeNewInputValue()
		answer = sol.CheckCorrect(count, sel, number)
	}
}
