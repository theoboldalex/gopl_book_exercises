// 1.7: refactor to io.Copy(dst, src) to copy resp.Body to stdout
// 1.8: check for protocol prefix in args and prepend if necessary using strings.HasPrefix
// 1.9: modify the program to also print the response status code

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const protocol = "https://"

func main() {
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, protocol) {
			url = protocol + url
		}

		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		io.Copy(os.Stdout, resp.Body)
		io.Copy(os.Stdout, strings.NewReader(resp.Status))
		resp.Body.Close()
	}
}
