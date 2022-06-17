package main

import "fmt"

func main() {
	//for condicional incrementando
	for i := 0; i <= 10; i++ {
		fmt.Println((i))
	}

	fmt.Printf("\n")

	//for condicional decrementando
	for i := 10; i >= 0; i-- {
		fmt.Println((i))
	}

	fmt.Printf("\n")

	//for while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	//counter Forever
	/* 	counterForever := 0
	   	for {
	   		fmt.Println(counterForever)
	   		counterForever++
	   	} */
}
