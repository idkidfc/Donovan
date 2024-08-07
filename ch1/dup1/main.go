package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		if in.Text() == "break" {
			break
		}
		counts[in.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
