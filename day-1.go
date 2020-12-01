package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day-1.input", "Relative file path")
var partB = flag.Bool("partB", false, "Whether to us the part B logic")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)

	split := strings.Split(contents, "\n")
	seen := make([]int, len(split)-1)
	for i, s := range split {
		if i == len(split)-1 {
			break
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Failed to parse %s\n", s)
			break
		}
		seen[i] = n
	}
	for i, n := range seen {
		for j, m := range seen {
			if i >= j {
				continue
			}
			if n+m == 2020 {
				fmt.Println(n * m)
				return
			}
		}
	}
}
