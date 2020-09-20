package main

import (
	"fmt"
)

func main(){
	var x int = 42
	fmt.Printf("%d\t%#x\t%b\t\n",x,x,x)
	b := x << 1
	fmt.Printf("%d\t%#x\t%b\t",b,b,b)
}

