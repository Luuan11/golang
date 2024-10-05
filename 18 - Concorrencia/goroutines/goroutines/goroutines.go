package main

import (
	"fmt"
	"time"
)

func main() {
	//CONCORRENCIA

	go write("ola mundo bonito")
	write("Amo golang")
}

func write(text string){
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}