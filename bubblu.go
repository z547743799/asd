package main

import (
	"fmt"
)

func main() {

	var list = []int{234, 354, 42, 23, 20, 111}
	ch1 := make(chan []int)
	ch2 := make(chan bool)
	go func(list []int) {
		ch1 <- list
		for i := 0; i < len(list); i++ {
			for j := i; j < len(list)-1; j++ {
				if list[i] > list[j+1] {
					list[i] = list[i] ^ list[j+1]
					list[j+1] = list[i] ^ list[j+1]
					list[i] = list[i] ^ list[j+1]
				}
			}
		}
		close(ch1)
	}(list)
	go func(list chan []int) {
		a := <-list
		fmt.Println(a)
		ch2 <- true
	}(ch1)
	<-ch2
}
