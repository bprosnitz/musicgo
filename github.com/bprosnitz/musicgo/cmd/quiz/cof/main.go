package main

import (
	"fmt"
	"math/rand"

	. "github.com/bprosnitz/musicgo"
	"github.com/bprosnitz/musicgo/intervals"
)

func randomNote() Note {
	return Note(rand.Intn(12))
}

func fifth(n Note) Note {
	return n.Interval(intervals.Fifth)
}

func getResponse() (Note, error) {
	var answer string
	if _, err := fmt.Scanf("%s\n", &answer); err != nil {
		return 0, fmt.Errorf("Badly formatted answer: %v", err)
	}
	a, err := ParseNote(answer)
	if err != nil {
		return 0, fmt.Errorf("Error parsing note: %v", err)
	}
	return a, nil
}

func ask() {
	n := randomNote()
	f := fifth(n)

	var answer Note
	if rand.Intn(2) == 0 {
		fmt.Printf("What is a fifth above %v? ", n)
		answer = f

	} else {
		fmt.Printf("What is a fifth below %v? ", f)
		answer = n
	}
	a, err := getResponse()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	if answer == a {
		fmt.Printf("Correct!\n")
	} else {
		fmt.Printf("Sorry, the answer was: %v\n", answer)
	}
}

func main() {
	for {
		ask()
	}
}
