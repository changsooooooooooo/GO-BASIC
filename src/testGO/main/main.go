package main

import (
	"fmt"
	"testGO/calc"
)

func main(){
	fmt.Println("Hello Go Programming")
	fmt.Println(calc.Add(3, 4))
	fmt.Println(calc.Sub(2, 1))
	fmt.Print(calc.Multiply(2, 3))
}
