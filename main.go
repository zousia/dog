package main

import (
	"fmt"

	"github.com/zousia/dog"
)

type doggy struct {
	name string
	age  int
}

func main() {
	fido := doggy{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
}
