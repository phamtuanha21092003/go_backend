package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormatStr(), name)
	return message, nil
}

func randomFormatStr() string {
	formats := []string{
		"Hi, %v",
		"Hello %v",
		"Welcome %v",
	}
	return formats[rand.Intn(len(formats))]
}

func main() {
	fmt.Println("hello")
}
