package main

import (
	"github.com/sidav/golibrl/console"
	"strconv"
)

func inputIntValue(prompt string) int {
	inputString := ""
	for {
		drawInputPrompt(prompt, inputString)
		key := console.ReadKey()
		if key == "ENTER" {
			res, _ := strconv.Atoi(inputString)
			return res
		}
		if key == "BACKSPACE" && len(inputString) > 0 {
			inputString = inputString[:len(inputString) - 1]
		}
		if len(key) < 2 && key != " " {
			inputString += key
		}
	}
	return 0
}

func inputStringValue(prompt string) string {
	inputString := ""
	for {
		drawInputPrompt(prompt, inputString)
		key := console.ReadKey()
		if key == "ENTER" {
			return inputString
		}
		if key == "BACKSPACE" && len(inputString) > 0 {
			inputString = inputString[:len(inputString) - 1]
		}
		inputString += key
	}
	return ""
}

func drawInputPrompt(prompt, input string) {
	console.Clear_console()
	console.SetBgColor(console.BEIGE)
	defer console.SetBgColor(console.BLACK)
	console.SetFgColor(console.BLACK)
	_, ch := console.GetConsoleSize()
	console.PutString(prompt, 0, ch-2)
	console.PutString(">" + input + "_", 0, ch-1)
	console.Flush_console()
}
