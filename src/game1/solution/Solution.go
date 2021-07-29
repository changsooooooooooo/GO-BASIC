package solution

import (
	"fmt"
	"math/rand"
	"time"
)

func MakeRandom() int {
	rand.Seed(time.Now().UnixNano())
	randNumber := rand.Intn(100)
	return randNumber
}

func MakeNewInputValue() int {
	var s int
	_, _ = fmt.Scanln(&s)
	return s
}

func CheckCorrect(count int, selNumber int, number int) int {
	fmt.Println(number)
	if selNumber == number {
		fmt.Printf("축하합니다. 숫자를 맞추셨습니다. 시도횟수: %d\n", count)
		return count
	}
	if selNumber < number {
		fmt.Println("입력하신 숫자가 더 작습니다.")
		return count + 1
	}
	fmt.Println("입력하신 숫자가 더 큽니다.")
	return count + 1
}
