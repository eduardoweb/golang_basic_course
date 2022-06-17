package main

import (
	"fmt"
)

func main() {
	//Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415
	fmt.Println("pi:", pi2)
	fmt.Println("pi2:", pi2)

	//Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int
	fmt.Println(base, altura, area)

	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	//Area del cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

	x := 10
	y := 50

	//OPERACIONES ARITMETICAS
	//Suma
	result := x + y
	fmt.Println("Suma", result)

	//Resta
	result = y - x
	fmt.Println("Resta", result)

	//Multiplicacion
	result = y * x
	fmt.Println("Multiplicacion", result)

	//Division
	result = y / x
	fmt.Println("Division", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo", result)

	//Incremental
	x++
	fmt.Println("Incremental", x)

	//Decremental
	x--
	fmt.Println("Incremental", x)

	//Retos
	// Calculo del area del Rectangulo, Trapecio, Circulo
	baseR := 10
	alturaR := 50
	areaRectangulo := baseR * alturaR
	fmt.Println("Area Rectangulo: ", areaRectangulo)

	baseA := 10
	baseB := 20
	alturaT := 30
	areaTrapecio := (baseA + baseB) * alturaT / 2
	fmt.Println("Area Trapecio: ", areaTrapecio)

	radio := 20.0
	var areaCirculo float64 = (radio * radio) * pi
	fmt.Println("Area Circulo: ", areaCirculo)

}
