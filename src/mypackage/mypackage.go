package mypackage

import "fmt"

//CarPublic with access public
type CarPublic struct {
	Brand string
	Year  int
}

func PrintMessage(text string) {
	fmt.Println(text)
}
