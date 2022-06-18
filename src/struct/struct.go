package main

import (
	pk "curso_golang_platzi/src/mypackage"
	"fmt"
)

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	//Otra manera
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)

	//LLamando un struct en otra pagina
	var myCar2 pk.CarPublic
	myCar2.Brand = "Lambo"
	myCar2.Year = 2020
	fmt.Println(myCar2)

	pk.PrintMessage("Hola Edu")
}
