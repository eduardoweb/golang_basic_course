package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	//Recorer el maps
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar un valor. la variable declarada "ok" nos indica si existe la llave que estamos buscando ejemplo: si existe m["jose"]. Si no existe por ejemplo: m["JOSER"] (coloque jose en mayuscula)
	//Retornara false, si existe retorna true
	value, ok := m["jose"]
	fmt.Println(value, ok)
}
