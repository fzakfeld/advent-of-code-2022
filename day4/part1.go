package main

import (
	"os"
	"strconv"
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
	pairs := strings.Split(st, "\n")

	total := 0

	for _, pair := range pairs {
		pair1 := strings.Split(strings.Split(pair, ",")[0], "-")
		pair2 := strings.Split(strings.Split(pair, ",")[1], "-")
		pair1_left, err := strconv.Atoi(pair1[0])
		pair1_right, err := strconv.Atoi(pair1[1])
		pair2_left, err := strconv.Atoi(pair2[0])
		pair2_right, err := strconv.Atoi(pair2[1])
		check(err)

		if (pair1_left <= pair2_left && pair1_right >= pair2_right) || (pair2_left <= pair1_left && pair2_right >= pair1_right) {
			total++
		}
	}

	println(total)
}
