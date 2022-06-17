package main

import "fmt"

func main() {
	//Array
	//Los array son inmutablers, ers decir, no puedes agregar elementos
	//len: Me indica la cantidad elementos del array
	//cap: me indica la capacidad maxima del arraay
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	//Slice: No defines el tamaño del slice
	//Son mutables, puedes agregar elememntos
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	//Metodos en el Slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])  //Imprime desde el primero hasta la posicion 3, el numero despues de los : es exclusivo, es decir,  no se va a imprimir
	fmt.Println(slice[2:4]) //Imprime un raggo el primer numero antes de los : es inclusivo, el segundo es exclusivo
	fmt.Print(slice[4:])    //Imprime desde la posicion 4 en adelante

	//Append
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append nueva lista
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...) //los ... indican que el slice nuevo lo va a descomprimir y en lugar de agregar una nueva lista quitara cada uno de los elementos y los agregar de forma independiente
	fmt.Println(slice)

	slice2 := []int{}
	for i := 0; i <= 20; i++ {
		slice2 = append(slice2, i)
	}
	fmt.Println(slice2)

}
