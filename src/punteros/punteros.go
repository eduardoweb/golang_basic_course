package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}
func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func main() {
	a := 50
	b := &a // El & significa buscar la direccion en memoria

	fmt.Println(b)
	fmt.Println(*b) // el * significa buscar el valor en esa direccion de memoria

	*b = 100       // Esto significa que estoy modificando el valor en esa direccion de memoria
	fmt.Println(a) // Al imprimir a nos va a dar 100 ya que a y *b estan apuntando a la misma direccion de memoria

	myPc := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPc)

	myPc.ping()

	fmt.Println(myPc)
	myPc.duplicateRAM()

	fmt.Println(myPc)
	myPc.duplicateRAM()

	fmt.Println(myPc)

}
