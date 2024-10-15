package main

func main(){

}

func fibonacci(posicao uint) uint{
	if posicao <= 1{
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}