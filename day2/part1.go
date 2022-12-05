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
	rounds := strings.Split(st, "\n")

	items := map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
	}

	points := map[string]int64{
		"rock":     1,
		"paper":    2,
		"scissors": 3,
		"won":      6,
		"draw":     3,
		"lost":     0,
	}

	evalMap := map[string]string{
		"rock":     "scissors",
		"paper":    "rock",
		"scissors": "paper",
	}

	var total int64 = 0

	for _, x := range rounds {
		val := strings.Split(x, " ")
		they := items[val[0]]
		me := items[val[1]]
		var state string
		if they == me {
			state = "draw"
		} else if evalMap[me] == they {
			state = "won"
		} else {
			state = "lost"
		}
		total += points[state]
		total += points[me]
	}

	println(total)

}
