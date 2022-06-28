package main

import (
	"fmt"
)

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "Mensaje1"
	c <- "Mensaje2"

	fmt.Print(len(c), cap(c))

	//Range y Close
	close(c) //Le indica al runtaime de Go que va a cerrar el canal

	for message := range c { //Rang esta perfecto para interactuar con cada uno de los mensajes dentro del canal, no importa si esta cerrado
		fmt.Println(message)
	}

	//Select
	email1 := make(chan string)
	email2 := make(chan string)
	go message("mensaje1", email1)
	go message("mensaje2", email2)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)

		}
	}
}
