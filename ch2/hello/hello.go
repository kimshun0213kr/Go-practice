package main

import (
	"fmt"
	"example.com/greetings"
)

func main(){
	message := greetings.Hello("Beleth")
	fmt.Println(message)
}