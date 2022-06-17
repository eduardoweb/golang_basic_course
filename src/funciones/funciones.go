package main

import "fmt"

func normalFuncion(message string) {
	fmt.Println(message)

}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFuncion("Hola Mundo")
	tripleArgument(1, 2, "hola")

	value := returnValue(2)
	fmt.Println("Value: ", value)

	value1, value2 := doubleReturn(2)
	fmt.Println("Value1 y Value2:", value1, value2)

	value3, _ := doubleReturn(2)
	fmt.Println("Value3:", value3)
}
