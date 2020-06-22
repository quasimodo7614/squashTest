package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("this is my test")
	for {
		time.Sleep(2*time.Second)
		fmt.Println("i had rested for 2 seconds")
	}
}
