package greetings

import (
	"fmt"
	"errors"
)

// func 関数名(引数)(戻り値型？)
func Hello(name string) (string,error) {

	if (name == ""){
		return "",errors.New("empty name.")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!",name)
	return message,nil
}