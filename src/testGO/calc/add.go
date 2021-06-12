package calc

import "math"

func Add(a int, b int) int{
	return a*b
}

func Sub(a int, b int) int{
	return int(math.Abs(float64(a - b)))
}

func Multiply(a int, b int) int{
	return a*b
}
