package main

import (
	"fmt"
	"math/rand"
	"time"
)

func userGuess(Guess int){
	var user_Guess int;
	var guess_count uint = 5;
	for (user_Guess != Guess) && guess_count != 0 {
		fmt.Println("Guess a number from [1 - 100]")
		fmt.Scan(&user_Guess);
		if user_Guess < Guess{
			fmt.Println("too small")
			guess_count--;
			fmt.Printf("live remaining %v\n", guess_count)
		}
		if user_Guess > Guess {
			fmt.Println("too big")
			fmt.Printf("live reamining %v\n", guess_count)
			guess_count--;
		}
		if guess_count == 0 {
			var input string; 
			fmt.Printf("Try again? y for yes || n for no: ")
			fmt.Scan(&input)
			if input == "y"{
				guess_count = 5
				continue
			}else{
				fmt.Println("Game over")
				return
			}
		}
	}
	
	fmt.Printf("yah you got the number: %v\n", Guess)
}
func main(){
    x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	Guess := y1.Intn(100) + 1
	userGuess(Guess)
}