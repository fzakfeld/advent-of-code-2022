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
	}

	outcomes := map[string]string{
		"X": "loose",
		"Y": "draw",
		"Z": "win",
	}

	points := map[string]int64{
		"rock":     1,
		"paper":    2,
		"scissors": 3,
		"win":      6,
		"draw":     3,
		"loose":    0,
	}

	evalMap := map[string][]string{
		// they: me win, me draw, me loose
		"rock":     {"paper", "rock", "scissors"},
		"paper":    {"scissors", "paper", "rock"},
		"scissors": {"rock", "scissors", "paper"},
	}

	var total int64 = 0

	for _, x := range rounds {
		val := strings.Split(x, " ")
		they := items[val[0]]
		outcome := outcomes[val[1]]
		println("they " + they)
		println("outcome" + outcome)
		var me string
		if outcome == "win" {
			me = evalMap[they][0]
		} else if outcome == "draw" {
			me = evalMap[they][1]
		} else {
			me = evalMap[they][2]
		}
		total += points[me]
		total += points[outcome]
	}

	println(total)

}
