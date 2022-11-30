package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

func main() {
	sep := " "
	src, err := readInput()
	if err != nil {
		fmt.Print(0)
		return
	}

	words := strings.Split(src, sep)
	fmt.Print(len(words))
}

func readInput() (string, error) {
	flag.Parse()
	src := strings.Join(flag.Args(), "")
	if src == "" {
		return src, errors.New("missing string to count")
	}

	return src, nil
}
