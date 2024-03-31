package main

import (
	"fmt"
	"math/rand"
	"time"
)

func movePolice(result chan int, pos chan int, n int, m int) {
	var x, y, move, r int
	x = 1
	y = n
	var bench = int(float32(n) / float32(n+m) * 100)

	fmt.Printf("Starting position -- Police: (%d, %d)\n", x, y)

	end := false
	for !end {
		move = rand.Intn(100)

		if move <= bench { // vertical
			move = rand.Intn(3)
			if move <= 1 { // south
				if (y - 1) >= 1 {
					y = y - 1
				} else {
					y = y + 1
				}
			} else { // north
				if (y + 1) <= n {
					y = y + 1
				} else {
					y = y - 1
				}
			}
		} else { // horizontal
			move = rand.Intn(3)
			if move <= 1 { // east
				if (x + 1) <= m {
					x = x + 1
				} else {
					x = x - 1
				}
			} else { // west
				if (x - 1) >= 1 {
					x = x - 1
				} else {
					x = x + 1
				}
			}
		}

		pos <- x
		pos <- y

		r = <-result
		if r != 0 {
			end = true
		}
	}
}

func moveThief(result chan int, pos chan int, n int, m int) {
	var x, y, move, r int
	x = m
	y = 1
	var bench = int(float32(n) / float32(n+m) * 100)

	fmt.Printf("Starting position -- Thief: (%d, %d)\n", x, y)

	end := false
	for !end {
		move = rand.Intn(100)

		if move <= bench { // vertical
			move = rand.Intn(3)
			if move <= 0 { // south
				if (y - 1) >= 1 {
					y = y - 1
				} else {
					y = y + 1
				}
			} else { // north
				if (y + 1) <= n {
					y = y + 1
				} else {
					y = y - 1
				}
			}
		} else { // horizontal
			move = rand.Intn(3)
			if move <= 0 { // east
				if (x + 1) <= m {
					x = x + 1
				} else {
					x = x - 1
				}
			} else { // west
				if (x - 1) >= 1 {
					x = x - 1
				} else {
					x = x + 1
				}
			}
		}

		pos <- x
		pos <- y

		r = <-result
		if r != 0 {
			end = true
		}
	}
}

// Labelling for each player (values passed onto res channels)
// 1) The game ends.
// 2) The game continues.

// Labelling for end-game output (value held in res)
// 1) The Police caught the Thief and won the game.
// 2) The Thief escaped and won the game.
// 3) The Police ran out of moves and the Thief won the game.
// 4) The game ends in a tie.

func controller(result1 chan int, pos1 chan int, result2 chan int, pos2 chan int, n int, m int, s int) {
	var x1, y1, x2, y2 int
	var x2_old, y2_old int

	x2_old = m
	y2_old = 1

	i := 1
	res := -1

	end := false
	for !end {
		x1 = <-pos1
		y1 = <-pos1

		x2 = <-pos2
		y2 = <-pos2

		fmt.Printf("\n---Round %d---\n", i)
		fmt.Printf("Police at position (%d, %d).\n", x1, y1)
		fmt.Printf("Thief at position (%d, %d).\n", x2, y2)

		// End-game logic
		if (x1 == x2_old && y1 == y2_old) || (x1 == x2 && y1 == y2) {
			if x1 == 1 && y1 == n { // tie
				result1 <- 1
				result2 <- 1
				res = 4
				end = true
			} else { // police wins
				result1 <- 1
				result2 <- 1
				res = 1
				end = true
			}
		} else {
			if x2 == 1 && y2 == n { // thief wins
				result1 <- 1
				result2 <- 1
				res = 2
				end = true
			} else if i >= s { // thief wins
				result1 <- 1
				result2 <- 1
				res = 3
				end = true
			} else { // game continues
				result1 <- 0
				result2 <- 0
			}
		}

		x2_old = x2
		y2_old = y2

		i = i + 1
	}

	if res == 1 {
		fmt.Printf("\nThe Police caught the Thief at (%d, %d) and won the game.\n", x1, y1)
	} else if res == 2 {
		fmt.Println("\nThe Thief escaped and won the game.")
	} else if res == 3 {
		fmt.Println("\nThe Police ran out of moves and the Thief won the game.")
	} else if res == 4 {
		fmt.Println("\nThe game ends in a tie.")
	} else {
		fmt.Println("We have an issue.")
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
