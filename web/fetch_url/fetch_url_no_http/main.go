package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const http_prefix = "http://"

var valid_prefixes = []string{"http://", "https://"}

func main() {
	for _, url := range os.Args[1:] {
		for n, prefix := range valid_prefixes {
			if strings.HasPrefix(url, prefix) {
				break
			}
			if n == len(valid_prefixes)-1 {
				url = http_prefix + url
			}
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		resp.Body.Close()

	}
}
