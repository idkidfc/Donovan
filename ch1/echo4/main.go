package main

import (
	"fmt"
	"os"
)

func main() {
	s := ""
	for i, arg := range os.Args[1:] {
		s += fmt.Sprintf("index:%d, arg:%s\n", i, arg)
	}
	fmt.Println(s)
}
