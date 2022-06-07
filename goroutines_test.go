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
