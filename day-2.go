package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day-2.input", "Relative file path")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n")
	var totalValid, totalValidB int

	for _, s := range split {
		splitSeen := strings.Split(s, " ")
		amounts := strings.Split(splitSeen[0], "-")
		from, err := strconv.Atoi(amounts[0])
		to, err := strconv.Atoi(amounts[1])
		if err != nil {
			fmt.Printf("Failed to parse %s\n", s)
			break
		}
		letter := strings.TrimSuffix(splitSeen[1], ":")
		password := splitSeen[2]

		t := strings.Count(password, letter)
		if t >= from && t <= to {
			totalValid++
		}

		a := string(password[from-1])
		b := string(password[to-1])
		if a != b && (a == letter || b == letter) {
			totalValidB++
		}
	}
	fmt.Println(totalValid)
	fmt.Println(totalValidB)

}
