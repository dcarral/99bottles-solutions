package main

import (
	"strconv"
	"strings"
)

func main() {
	Sing()
}

func Sing() string {
	return Verses(99, 1)
}

func Verses(start int, end int) string {
	var verses string
	for i := start; i >= end; i-- {
		verses = verses + Verse(i) + "\n"
	}
	return verses
}

func Verse(number int) string {
	var current_quantity string
	var next_quantity string
	action := "Take one down and pass it around, "
	switch number {
	case 2:
		current_quantity = "2 bottles"
		next_quantity = "1 bottle"
	case 1:
		current_quantity = "1 bottle"
		next_quantity = "no more bottles"
	case 0:
		current_quantity = "no more bottles"
		next_quantity = strconv.Itoa(99) + " bottles"
		action = "Go to the store and buy some more, "
	default:
		current_quantity = strconv.Itoa(number) + " bottles"
		next_quantity = strconv.Itoa(number-1) + " bottles"
	}

	return Capitalize(current_quantity) + " of beer on the wall, " + current_quantity + " of beer.\n" + action + next_quantity + " of beer on the wall.\n"
}

func Capitalize(title string) string {
	words := strings.Fields(title)
	if _, err := strconv.Atoi(words[0]); err == nil {
		return title
	}

	words[0] = strings.Title(words[0])
	return strings.Join(words, " ")
}
