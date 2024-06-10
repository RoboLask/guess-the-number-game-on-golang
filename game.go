package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("guess the number from 1 to 10")
	time.Sleep(1 * time.Second)
	a := rand.Intn(11)
	var r int
	for {
		fmt.Scan(&r)
		if r == a {
			fmt.Println("You guessed it!")
			time.Sleep(1 * time.Second)
			break
		} else if r < a {
			fmt.Println("the number is greater!")
		} else {
			fmt.Println("your number is less!")
		}
	}
}
