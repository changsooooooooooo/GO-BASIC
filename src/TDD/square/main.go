package square

import "fmt"

func Square(x int) int {
	return x * x
}

func main() {
	fmt.Printf("9*9=%d\n", Square(9))
}
