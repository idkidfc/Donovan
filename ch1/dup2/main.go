package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
		return
	}
	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
		f, err := os.Open(arg)
		fmt.Printf("%+v", f)
		defer f.Close()
		if err != nil {
			log.Println("cannot open a file")
			continue
		}
		countLines(f, counts)
	}

	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}

func countLines(f io.Reader, data map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		data[input.Text()]++
	}
}
