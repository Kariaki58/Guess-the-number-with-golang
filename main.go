package main

import (
	"fmt"
	"math/rand"
	"time"
)
func role_dice() (int, int){
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	randn_1 := y1.Intn(6) + 1
	randn_2 := y1.Intn(6) + 1
	
	return randn_1, randn_2
}
func main(){
	var rand_1 int;
	var randn_2 int;
	for i := 0; i < 1000; i++{
		time.Sleep(time.Second)
		rand_1, randn_2 = role_dice()
		fmt.Printf("(%v, %v)\n", rand_1, randn_2)
	}
}