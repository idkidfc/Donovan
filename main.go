package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		some()
	}
}

func some() {

	st := time.Now()
	defer func() {
		fmt.Println("defer")
		fmt.Println(time.Now().Sub(st))
	}()

	return
}
