package main

import (
	"fmt"
)

func main() {

	/* 	************* COMANDOS UTILIZADOS EN EL CURSO DE GO BASICO ***************************
	   	1) go run src/struct/struct.go = Corre el proyecto o pagina donde estas posicionado

	   	2) go mod init = Crea el modulo. Anteriormente en GO se tenía que trabajar el código dentro del GOPATH. Este GOPATH es una de las variables de entorno que declaramos al principio del curso (export GOPATH=$HOME/go) y no es más que la ruta a nuestro entorno de trabajo en donde encontraremos/crearemos la carpeta source que es donde trabajamos todo nuestro código.
	   	Afortunadamente desde la versión 1.11 podemos realizar nuestros proyectos de Go desde donde queramos gracias a los módulos.
	   	Veamos un ejemplo. Cree un proyecto llamando 0.0-fundamentals en la siguiente ruta /home/est14/my-own-path/15-go-course/0.0-fundamentals como se puede apreciar estoy fuera de mi GOPATH (/home/est14/go/source).
	   	Ahora cree mi go.mod aquí, en la raíz de 0.0-fundamentals con el siguiente comando. go mod init est14/0.0-fundamentals  Esto me creo el archivo go.mod con el siguiente contenido :
	   	module est14/0.0-fundamentals
	   		go 1.16  // Esta es la version de Go
	   	Luego cree un archivo llamado main.go, al mismo nivel de este cree un directorio llamado greetings, dentro de greetings cree un archivo llamado greetings.go y finalmente cree una función llamada Hola.
	   	Mi estructura quedo asi:
	   	├── go.mod
	   	├── greetings
	   	│ └── greetings.go // Aquí cree mi funcion Hola
	   	└── main.go
	   	Ahora mi intención era traer la función llamada Hola que estaba en greetings.go a mi archivo main.go y eso lo hice importando mi paquete “est14/0.0-fundamentals/grettings”
	   	//Este es el contenido de mi archivo main.go
	   	package main

	   	import (
	   		"est14/0.0-fundamentals/grettings"
	   	)

	   	func main() {
	   		greetings.Hola()
	   	}
	   	Y listo finalmente pude importar mi paquete y utilizar mi función Hola.
	   	Hay mucho detrás de todo esto y sería bueno detenerse un poco aquí ir por un helado luego leer un poco de documentación y seguir .

	   	3) go build src/mypackage/mypackage.go = Este comando crea y compila el programa
	*/

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
