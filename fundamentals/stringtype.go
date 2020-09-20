package main

import (
	"fmt"
)

func main(){
	s := "Hello Arda"

	fmt.Println(s)
	fmt.Println(len(s))

	fmt.Printf("%T\n", s)
	fmt.Println([]byte(s))

}

