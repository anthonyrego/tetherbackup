package main

import "fmt"

const black = "30;1"
const red = "31;1"
const green = "32;1"
const yellow = "33;1"
const blue = "34;1"
const magenta = "35;1"
const cyan = "36;1"
const white = "37;1"
const clear = "0"

func colorizeString(text string, color string) string {
	return fmt.Sprintf("\x1b[%sm%s\x1b[0m", color, text)
}

func clearScreen() {
	fmt.Printf("\033[2J\033[;H")
}
