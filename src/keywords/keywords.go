package main

import "fmt"

func main() {

	/* Defer: este keywords se utiliza para indicar que esta linea de codigo se ejecutara al final de una funcion, es deicr, no importa
	donde la escribas, al colocarle defer go ejecutara esa linea al final de la funcion */
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//Continue y Break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		//Break
		if i == 8 {
			fmt.Println("break")
			break
		}
	}
}
