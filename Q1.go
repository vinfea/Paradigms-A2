package main

import (
	"fmt"
	"math/rand"
	"time"
)

func movePolice(result chan int, pos chan int) {
	fmt.Println("In Police Go Routine.")
}

func moveThief(result chan int, pos chan int) {
	fmt.Println("In Thief Go Routine.")
}

func controller(result1 chan int, pos1 chan int, result2 chan int, pos2 chan int, n int, m int, s int) {
	fmt.Println("In Controller Go Routine.")
}

func main() {
	var n, m, max1, max2, s int

	n = rand.Intn(191) + 10
	m = rand.Intn(191) + 10
	max1 = 2 * max(n, m)
	max2 = 10 * max(n, m)
	s = rand.Intn(max2-max1+1) + max1

	result1 := make(chan int)
	result2 := make(chan int)
	pos1 := make(chan int)
	pos2 := make(chan int)

	fmt.Printf("The 黑猫警长 Game begins with a %d x %d grid.\n", n, m)
	fmt.Printf("The number of moves for the Police to catch the Thief is %d.\n", s)

	go movePolice(result1, pos1)
	go moveThief(result2, pos2)
	go controller(result1, pos1, result2, pos2, n, m, s)

	time.Sleep(2 * time.Second)
	fmt.Println("The 黑猫警长 Game is finished!")
}
