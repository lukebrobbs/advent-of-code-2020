package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day-2.input", "Relative file path")
var partB = flag.Bool("partB", false, "Whether to us the part B logic")

type password struct {
	amountFrom, amountTo int
	letter, password     string
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n")
	passwords := make([]password, len(split))
	for i, s := range split {
		splitSeen := strings.Split(s, " ")
		amounts := strings.Split(splitSeen[0], "-")
		from, err := strconv.Atoi(amounts[0])
		if err != nil {
			fmt.Printf("Failed to parse %s\n", s)
			break
		}
		to, err := strconv.Atoi(amounts[1])
		if err != nil {
			fmt.Printf("Failed to parse %s\n", s)
			break
		}
		passwords[i] = password{
			amountFrom: from,
			amountTo:   to,
			letter:     strings.TrimSuffix(splitSeen[1], ":"),
			password:   splitSeen[2],
		}
	}
	var totalValid int
	for _, p := range passwords {
		a := strings.Count(p.password, p.letter)
		if a >= p.amountFrom && a <= p.amountTo {
			totalValid++
		}
	}
	fmt.Println(totalValid)
	if *partB {
		var totalValidB int
		for _, p := range passwords {
			a := string(p.password[p.amountFrom-1])
			b := string(p.password[p.amountTo-1])
			if a != b && (a == p.letter || b == p.letter) {
				totalValidB++
			}
		}
		fmt.Println(totalValidB)
	}
}
