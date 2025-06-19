package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

func Hello(name string) (string,error) {

	if (name == ""){
		return "",errors.New("empty name.")
	}
	// ↓では%vを含む形で文字列が関数から帰ってきていて、それにnameが代入されている？
	message := fmt.Sprintf(randomFormat(),name)
	return message,nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. welcome!",
		"Great to see you, %v!",
		"こんにちは%v!お待ちしてました！",
	}
	return formats[rand.Intn(len(formats))]
}