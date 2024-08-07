package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	counts := make(map[stgring]int)
	for _, fn := range os.Args[1:] {
		data, err := ioutil.ReadFile(fn)
		if err != nil {
			log.Printf("cannot open a file, err:%v\n", err)
			return
		}
		for _, line := range data {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}
