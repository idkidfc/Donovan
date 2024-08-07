package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Counter struct {
	Mu    *sync.RWMutex
	Count int64
}

func (c *Counter) handler(w http.ResponseWriter, r *http.Request) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Count++
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
func (c *Counter) counter(w http.ResponseWriter, r *http.Request) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	fmt.Fprintf(w, "Count %d\n", c.Count)
}
func main() {
	c := &Counter{Mu: new(sync.RWMutex), Count: 0}

	http.HandleFunc("/", c.handler)
	http.HandleFunc("/count", c.counter)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
