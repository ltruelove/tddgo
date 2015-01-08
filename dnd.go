package main

import (
	"fmt"
)

func main() {
	myChar := NewCharacter()
	myChar.Name = "Whatever"
	fmt.Println(myChar.Name)
}
