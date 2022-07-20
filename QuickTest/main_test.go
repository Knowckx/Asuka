package main

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

var DefaultTimeoutLimit = 3 * time.Minute

func Test_Func1(t *testing.T) {
	out, err := Func1()
	if err != nil {
		fmt.Printf("%+v", err)
	}
	assert.Nil(t, err)
	fmt.Println(out)
}
func Func1() (int, error) {
	fmt.Printf("%v\n", DefaultTimeoutLimit)
	now := time.Now().UTC().Local()
	res := now.Format(time.RFC3339)
	fmt.Println(res)
	return 0, nil
}
