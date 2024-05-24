package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func(){
		write("Olá")
		waitGroup.Done()
	} ()

	go func(){
		write("Mundo")
		waitGroup.Done()
	} ()
}

func write(text string){
	for i := 0; i < 5; i++ {
		fmt.Print(text)
		time.Sleep(time.Second)
	}
}