package main

import (
	"fmt"
	"time"
)

func funcA(c chan int) {
	time.Sleep(time.Duration(2) * time.Second)
	c <- 18
}

func funcB(c chan int) {
	time.Sleep(time.Duration(3) * time.Second)
	c <- 96
}

func main() {
	c := make(chan int)
	go funcA(c)
	go funcB(c)
	num1 := <-c
	num2 := <-c
	fmt.Println(num1, num2)
}
