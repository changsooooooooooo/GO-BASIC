package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int){
	for n:=range ch{
		fmt.Println(n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func main(){
	var wg sync.WaitGroup

	ch := make(chan int, 2)

	wg.Add(1)
	go square(&wg, ch)

	for i:=0;i<10;i++{
		ch<-i*2
	}

	wg.Wait()
}
