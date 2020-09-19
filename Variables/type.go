package main

import (
	"fmt"
)

var y int = 25
func main(){
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	//y = "Arda" // because go lang is static type not dynamic type you dont change int to string or other type
	//fmt.Println(y)
}

