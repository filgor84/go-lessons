package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for fid, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
			fmt.Printf("%d - %s:\n", fid, arg)
			for line, n := range counts {
				if n > 1 {
					fmt.Printf("%d\t%s\n", n, line)
				}
			}
			counts = make(map[string]int)

		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
