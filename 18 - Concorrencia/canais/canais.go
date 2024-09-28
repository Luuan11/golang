package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go write("Hello world", canal)

	fmt.Println("after the function")

	for {
		msg, open := <-canal
		if !open {
			break
		}
		fmt.Println(msg)
	}
	fmt.Println("end of the program")

}

func write(text string, canal chan string) {
	// time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		canal <- text
		time.Sleep(time.Second)
	}

	close(canal)
}
