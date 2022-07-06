package main

import (
	"fmt"
	"sync"
	"testing"
)

func Test_Once(t *testing.T) {
	var once sync.Once
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceFunc)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
	fmt.Println(KeyCount)
}

var KeyCount int

func onceFunc() {
	KeyCount++
}
