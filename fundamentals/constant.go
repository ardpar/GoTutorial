package main

import (
	"fmt"
)

const (
	a int = 72
	b float64 = 42.78
	c string = "Arda Parlak"
)

func main(){
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
