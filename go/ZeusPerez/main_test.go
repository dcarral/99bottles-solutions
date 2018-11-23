package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func ReadFile() string {
	bytes, errors := ioutil.ReadFile("full_song.txt") // just pass the file name
	if errors != nil {
		fmt.Print(errors)
	}

	str := string(bytes) // convert content to a 'string'

	return str
}

func TestVerse99(test *testing.T) {
	expected_output := "99 bottles of beer on the wall, 99 bottles of beer.\nTake one down and pass it around, 98 bottles of beer on the wall.\n"
	output := Verse(99)
	if expected_output != output {
		test.Error("The verse is not the expected one")
	}
}

func TestVerse3(test *testing.T) {
	expected_output := "3 bottles of beer on the wall, 3 bottles of beer.\nTake one down and pass it around, 2 bottles of beer on the wall.\n"
	output := Verse(3)
	if expected_output != output {
		test.Error("The verse is not the expected one")
	}
}

func TestVerse2(test *testing.T) {
	expected_output := "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n"
	output := Verse(2)
	if expected_output != output {
		test.Error("The verse is not the expected one")
	}
}
func TestVerse1(test *testing.T) {
	expected_output := "1 bottle of beer on the wall, 1 bottle of beer.\nTake one down and pass it around, no more bottles of beer on the wall.\n"
	output := Verse(1)
	if expected_output != output {
		test.Error("The verse is not the expected one")
	}
}

func TestVerse0(test *testing.T) {
	expected_output := "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"
	output := Verse(0)
	if expected_output != output {
		test.Error("The verse is not the expected one")
	}
}

func TestSong(test *testing.T) {
	song := ReadFile()
	singedSong := Sing()
	if singedSong != song {
		test.Error("The singed song is not the expected one")
	}
}
