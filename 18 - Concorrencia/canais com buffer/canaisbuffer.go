package main

import "fmt"

func main() {
	canal := make(chan string, 2)

	canal <- "Olá Mundo!"
	canal <- "Olá Go"

	msg := <-canal
	fmt.Println(msg)
}