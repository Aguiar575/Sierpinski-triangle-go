package main

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

var grain = "*"

func main() {
	for {
		clearScreen()
		order := 5

		// Initialize the pyramid with one element containing the base unit surrounded by spaces
		fractalBody := []string{grain + strings.Repeat(" ", utf8.RuneCountInString(grain))}

		for order > 0 {
			// Calculate the number of spaces needed to center the base unit in the current level
			sp := strings.Repeat(" ", utf8.RuneCountInString(fractalBody[0])/2)

			top := cratePiramidSlice(fractalBody, sp)
			fractalBody = append(top, fractalBody...)

			printTheActualStateOfPiramid(fractalBody)

			order--
		}

		time.Sleep(500 * time.Millisecond)

		clearScreen()
	}
}

func cratePiramidSlice(t []string, sp string) []string {
	top := make([]string, len(t))

	setSpacesAroundTheGrain(t, top, sp)
	return top
}

func setSpacesAroundTheGrain(t []string, top []string, sp string) {
	for i, s := range t {
		top[i] = sp + s + sp
		t[i] += s
	}
}

func printTheActualStateOfPiramid(t []string) {
	// This helps to create the animation effect
	moveCursorToTopLeft()

	for _, r := range t {
		fmt.Println(r)
	}

	time.Sleep(500 * time.Millisecond)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func moveCursorToTopLeft() {
	fmt.Print("\033[H")
}
