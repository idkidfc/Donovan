package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	ch := make(chan string)

	wg := new(sync.WaitGroup)
	for _, url := range os.Args[1:] {
		wg.Add(1)
		go fetch(wg, url, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(wg *sync.WaitGroup, url string, ch chan<- string) {
	defer wg.Done()
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%+v", err)
		return
	}
	defer resp.Body.Close()

	ch2 := make(chan int64)
	go func() {

		fn, err := URLtoFileName(url)
		if err != nil {
			ch <- fmt.Sprintf("%+v", err)
			ch2 <- -1
		}

		n, err := ToFile(fn, resp.Body)
		if err != nil {
			ch <- fmt.Sprintf("%+v", err)
			ch2 <- -1
		}
		ch <- fmt.Sprintf("while writing to %s: %+v", url, err)
		ch2 <- n
	}()

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, <-ch2, url)
}

func ToFile(fileName string, src io.Reader) (int64, error) {
	f, err := os.Create(fileName)
	if err != nil {
		return -1, err
	}
	n, err := io.Copy(f, src)
	if err != nil {
		return -1, err
	}
	return n, nil
}

func URLtoFileName(url string) (string, error) {
	re := regexp.MustCompile(`https?://([^/]+)`)
	match := re.FindStringSubmatch(url)
	if len(match) <= 1 {
		return "", fmt.Errorf("URL length less than 1")
	}
	domain := match[1]
	formatted := strings.ReplaceAll(domain, ".", "_")
	return formatted, nil
}
