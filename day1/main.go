package main

import (
	"log"
	"os"
	"sort"
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
	st := string(data) + "\n"

	temp := strings.Split(st, "\n")

	currentFood := 0
	var foods []int

	for _, v := range temp {
		if v != "" {
			val, err := strconv.Atoi(v)
			check(err)
			currentFood += val
		} else {
			foods = append(foods, currentFood)
			currentFood = 0
		}
	}

	sort.Ints(foods)
	log.Println(foods[len(foods)-1] + foods[len(foods)-2] + foods[len(foods)-3])

}
