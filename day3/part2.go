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

	var contents [][][]string

	for i := 0; i < len(rucksacks); i += 3 {
		var teamContents [][]string
		teamContents = append(teamContents, strings.Split(rucksacks[i], ""))
		teamContents = append(teamContents, strings.Split(rucksacks[i+1], ""))
		teamContents = append(teamContents, strings.Split(rucksacks[i+2], ""))
		contents = append(contents, teamContents)
	}

	all := 0

	for _, team := range contents {
		firstBackpack := team[0]
		var cardItem string
		for _, item := range firstBackpack {
			isThere := true
			for _, rucksack := range team {
				check := checkBackback(item, rucksack)
				if check == false {
					isThere = false
				}
			}
			if isThere == true {
				cardItem = item
			}
		}
		code := []rune(cardItem)[0]
		if code > 90 {
			all += (int(code) - 96)
		} else {
			all += (int(code) - 38)
		}
	}
	println(all)
}

func checkBackback(wanted string, rucksack []string) bool {
	for _, item := range rucksack {
		if item == wanted {
			return true
		}
	}
	return false
}
