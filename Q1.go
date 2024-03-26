package main

import (
	"fmt"
	"math/rand"
	"time"
)

func movePolice(result chan int, pos chan int, n int, m int) {
	var x, y, move int
	x = 1
	y = n

	fmt.Printf("Police: (%d, %d)\n", x, y)

	// using for loop for the time being
	// before end-game logic is implemented
	for i := 0; i < 10; i++ {
		move = rand.Intn(4)

		if move == 0 { // north
			if (y + 1) <= n {
				y = y + 1
			} else {
				y = y - 1
			}
		} else if move == 1 { // east
			if (x + 1) <= m {
				x = x + 1
			} else {
				x = x - 1
			}
		} else if move == 2 { // south
			if (y - 1) >= 1 {
				y = y - 1
			} else {
				y = y + 1
			}
		} else { // west
			if (x - 1) >= 1 {
				x = x - 1
			} else {
				x = x + 1
			}
		}

		pos <- x
		pos <- y
	}
}

func moveThief(result chan int, pos chan int, n int, m int) {
	var x, y, move int
	x = m
	y = 1

	fmt.Printf("Thief: (%d, %d)\n", x, y)

	// using for loop for the time being
	// before end-game logic is implemented
	for i := 0; i < 10; i++ {
		move = rand.Intn(4)

		if move == 0 { // north
			if (y + 1) <= n {
				y = y + 1
			} else {
				y = y - 1
			}
		} else if move == 1 { // east
			if (x + 1) <= m {
				x = x + 1
			} else {
				x = x - 1
			}
		} else if move == 2 { // south
			if (y - 1) >= 1 {
				y = y - 1
			} else {
				y = y + 1
			}
		} else { // west
			if (x - 1) >= 1 {
				x = x - 1
			} else {
				x = x + 1
			}
		}

		pos <- x
		pos <- y
	}
}

func controller(result1 chan int, pos1 chan int, result2 chan int, pos2 chan int, n int, m int, s int) {
	var x1, y1, x2, y2 int

	// using for loop for the time being
	// before end-game logic is implemented
	for i := 0; i < 10; i++ {
		x1 = <-pos1
		y1 = <-pos1

		x2 = <-pos2
		y2 = <-pos2

		fmt.Printf("\n---Round %d---\n", i)
		fmt.Printf("Police at position (%d, %d).\n", x1, y1)
		fmt.Printf("Thief at position (%d, %d).\n", x2, y2)
	}
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
	pos1 := make(chan int, 2)
	pos2 := make(chan int, 2)

	fmt.Printf("The 黑猫警长 Game begins with a %d x %d grid.\n", m, n)
	fmt.Printf("The number of moves for the Police to catch the Thief is %d.\n", s)
	fmt.Println()

	go movePolice(result1, pos1, n, m)
	go moveThief(result2, pos2, n, m)
	go controller(result1, pos1, result2, pos2, n, m, s)

	time.Sleep(2 * time.Second)

	fmt.Println()
	fmt.Println("The 黑猫警长 Game is finished!")
}
