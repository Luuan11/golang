package main

import "fmt"

func main() {
	canal := make(chan string)

	canal <- "Olá Mundo!"

	msg := <-canal
	fmt.Println(msg)
}