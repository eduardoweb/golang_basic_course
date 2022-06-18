package main

import "fmt"

//Los Stringers sinver para customizar el output de un struct

type pc struct {
	ram   int
	brand string
	disk  int
}

func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPC.ram, myPC.disk, myPC.brand) //el Sprintf no imprime en consola, sino, que crea un string en pantalla con el resultado obtenido
}

func main() {
	myPc := pc{ram: 16, brand: "msi", disk: 100}

	fmt.Println(myPc)
}
