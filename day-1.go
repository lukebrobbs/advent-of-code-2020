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
	contains := make(map[int]bool)
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
		contains[n] = true
	}

partA:
	for _, n := range seen {
		if contains[2020-n] {
			fmt.Println(n * (2020 - n))
			break partA
		}
	}
	if *partB {
	partB:
		for i, m := range seen {
			for j, n := range seen {
				if i > j {
					continue
				}
				sumMN := n + m
				if sumMN >= 2020 {
					continue
				}
				if contains[2020-sumMN] {
					fmt.Println(m * n * (2020 - sumMN))
					break partB
				}

			}
		}
	}
}
