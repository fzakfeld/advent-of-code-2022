package main

import (
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile("./input.txt")
	check(err)
	st := string(data)
	rucksacks := strings.Split(st, "\n")

	all := 0

	for _, x := range rucksacks {
		size := len(x)
		var first []string
		var second []string
		for c := 0; c < size/2; c++ {
			first = append(first, string(x[c]))
		}
		for c := size / 2; c < size; c++ {
			second = append(second, string(x[c]))
		}
		var val string
		for _, c := range first {
			for _, b := range second {
				if c == b {
					val = c
				}
			}
		}
		code := []rune(val)[0]
		if code > 90 {
			all += (int(code) - 96)
		} else {
			all += (int(code) - 38)
		}
	}

	println(all)

}
