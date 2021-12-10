// Fetchall fetches URLs in parallel and reports their times and sizes.
// 1.10: Investigate caching by running fetchall twice. Modify fetchall to print its output to a file
// 1.11: How does the program behave if a web site doesn't respond?
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	f, err := os.Create("fetchall_output.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	for range os.Args[1:] {
		io.Copy(f, strings.NewReader(<-ch+"\n")) // receive from channel ch
	}

	io.Copy(f,
		strings.NewReader(fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())))
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
