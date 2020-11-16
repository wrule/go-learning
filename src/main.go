package main

import (
	"fmt"
	"time"
)

var num int = 0

func funcA() {
	for i := 0; i < 50000; i++ {
		num++
		fmt.Printf("A: %d\n", num)
		// time.Sleep(time.Duration(1) * time.Second)
	}
}

func funcB() {
	for i := 0; i < 50000; i++ {
		num++
		fmt.Printf("A: %d\n", num)
		// time.Sleep(time.Duration(1) * time.Second)
	}
}

func mainSleep() {
	for {
		time.Sleep(time.Duration(1) * time.Hour)
	}
}

func main() {
	// fmt.Println("你好世界")
	// go funcA()
	// go funcB()
	// fmt.Println("数据源是")
	// mainSleep()

	num := time.Now().UnixNano()
	fmt.Println(num)
	time.Sleep(time.Duration(1) * time.Second)
	num = time.Now().UnixNano()
	fmt.Println(num)
}
