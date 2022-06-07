package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func runHelloWorld() {
	fmt.Println("Hello world")
}

func TestCreateGoroutine(t *testing.T) {
	go runHelloWorld()
	fmt.Println("Ups")
	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Number", number)
}

func TestManyGoroutine(t *testing.T) {
	limit := 100000
	// limit := 20
	for number := 1; number <= limit; number++ {
		go DisplayNumber(number)
	}
	time.Sleep(5 * time.Second)
}
