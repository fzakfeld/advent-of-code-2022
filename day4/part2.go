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

		map1 := getUsedMap(pair1_left, pair1_right)
		map2 := getUsedMap(pair2_left, pair2_right)

		overlap := false

		for _, x := range map1 {
			for _, y := range map2 {
				if x == y {
					overlap = true
				}
			}
		}

		if overlap == true {
			total++
		}
	}

	println(total)
}

func getUsedMap(min int, max int) []int {
	var usedMap []int

	for i := min; i <= max; i++ {
		usedMap = append(usedMap, i)
	}

	return usedMap
}
