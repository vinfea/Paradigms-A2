package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func movePolice(result chan int, pos chan int, n int, m int, wg *sync.WaitGroup) {
	defer wg.Done()
	var x, y, move int
	x = 1
	y = n

	fmt.Printf("Police: (%d, %d)\n", x, y)

	// using for loop for the time being
	// before end-game logic is implemented
	//for i := 0; i < 10; i++ {
	//testing
	var end bool
	var r int
	end = false

	for !end {
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

		r = <-result //sends the integer from channel result, assigns it to variable r

		if r == -1 {
			end = true
		}
	}
	//}
}

func moveThief(result chan int, pos chan int, n int, m int, wg *sync.WaitGroup) {
	defer wg.Done()

	var x, y, move int
	x = m
	y = 1

	fmt.Printf("Thief: (%d, %d)\n", x, y)

	// using for loop for the time being
	// before end-game logic is implemented

	var end bool
	var r int
	end = false

	//for i := 0; i < 10; i++ {

	//testing
	for !end {

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

		r = <-result //sends the integer from channel result, assigns it to variable r

		if r == -1 {
			end = true
		}
	}
}

//conditions for win: if the police runs out of s
//the thief reaches the top left cell
//if the the police is qat the top left cell, and the police is at the top left cell,
//and the thief moves into that cell, the game ends in tie

func controller(result1 chan int, pos1 chan int, result2 chan int, pos2 chan int, n int, m int, s int, wg *sync.WaitGroup) {
	defer wg.Done()
	var x1, y1, x2, y2 int

	end := false
	//make the s the thing that its looping??

	// using for loop for the time being
	// before end-game logic is implemented
	//for i := 0; i < 10; i++ {
	var i int
	i = 0
	//testing
	for !end {
		x1 = <-pos1
		y1 = <-pos1

		x2 = <-pos2
		y2 = <-pos2

		fmt.Printf("\n---Round %d---\n", i)
		fmt.Printf("Police at position (%d, %d).\n", x1, y1)
		fmt.Printf("Thief at position (%d, %d).\n", x2, y2)

		if s == 0 {
			fmt.Println("Police ran out of moves. Thief wins!")

			result1 <- -1
			result2 <- -1

			end = true
		} else if x1 == 1 && y1 == n && x2 == 1 && y2 == 1 {
			fmt.Println("Police and Thief are at the top left cell. Game ends in a tie!")

			result1 <- -1
			result2 <- -1

			end = true
		} else if x2 == 1 && y2 == 1 {
			fmt.Println("Thief is at the top left cell. Thief wins!")

			result1 <- -1
			result2 <- -1

			end = true
		} else if x1 == x2 && y1 == y2 {
			fmt.Println("Police caught thief. Police wins!")

			result1 <- -1
			result2 <- -1

			end = true
		} else {
			s = s - 1
			i = i + 1

			result1 <- 0
			result2 <- 0

		}

	}
	//}
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

	var wg sync.WaitGroup
	wg.Add(3) // Wait for three goroutines to finish

	go movePolice(result1, pos1, n, m, &wg)
	go moveThief(result2, pos2, n, m, &wg)
	go controller(result1, pos1, result2, pos2, n, m, s, &wg)

	wg.Wait()

	//time.Sleep(2 * time.Second)

	fmt.Println()
	fmt.Println("The 黑猫警长 Game is finished!")
}
