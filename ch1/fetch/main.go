package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("cannot get req, err:%v", err)
			continue
		}
		defer resp.Body.Close()
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("cannot read req body, err:%v", err)
		}
		fmt.Printf("%d\n", len(b))
		for k, v := range resp.Header {
			fmt.Println(k, v)
		}
	}
}
