package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	valor1 := 1
	valor2 := 2
	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	//With and
	if valor1 == 1 && valor2 == 3 {
		fmt.Println("Es verdad")
	}

	//With or
	if valor1 == 1 || valor2 == 2 {
		fmt.Println("Es verdad, OR")
	}

	//Convertir texto a mnumero
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err) //Esto me arroja el error en consola y aplica un exit, es decir, detiene el codigo
	}
	fmt.Println("Value:", value)

	//Convertir texto a mnumero
	/* 	value2, err := strconv.Atoi("sdsdsdsd")
	   	if err != nil {
	   		log.Fatal(err) //Esto me arroja el error en consola y aplica un exit, es decir, detiene el codigo
	   	}
	   	fmt.Println("Value:", value2) */

	//Switchj con condicion
	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	//Switch sin condicion
	value3 := 200
	switch {
	case value3 > 100:
		fmt.Println("Es mayor a 100")
	case value3 < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condicion")
	}
}
