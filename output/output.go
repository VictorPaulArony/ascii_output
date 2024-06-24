package output

import (
	"fmt"
	"os"
	"strings"
)

// DisplayText displays the provided text along with content lines
func Display(input string, contentLines []string, banner string) {
	if input == "" {
		return
	}

	if input == "\\n" || input == "\n" {
		fmt.Println()
		return
	}
	// make newline and tab printable inthe terminal output
	input = strings.ReplaceAll(input, "\n", "\\n")
	input = strings.ReplaceAll(input, "\\t", "\t")

	wordslice := strings.Split(input, "\\n")

	var outputLines []string

	for _, word := range wordslice {
		if word == "" {
			outputLines = append(outputLines, "\n")
		} else {
			if English(word) {
				lines := PrintAsciiWord(word, contentLines)
				outputLines = append(outputLines, lines...)
			} else {
				outputLines = append(outputLines, "Invalid input: not accepted")
			}
		}
	}
	err := os.WriteFile(banner, []byte(strings.Join(outputLines, "\n")), 0o644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

// IsEnglish checks if a word contains only English alphabets
func English(words string) bool {
	for _, word := range words {
		if word < 32 || word > 126 {
			return false
		}
	}
	return true
}

// PrintWord prints a word if it exists in the content lines
func PrintAsciiWord(word string, contentLines []string) []string {
	linesOfSlice := make([]string, 9)

	for _, v := range word {
		for i := 1; i <= 9; i++ {
			linesOfSlice[i-1] += contentLines[int(v-32)*9+i]
		}
	}
	return linesOfSlice
	// fmt.Print(strings.Join(linesOfSlice, "\n"))
	// os.WriteFile(banner,data[strings.Join(linesOfSlice, "\n")]byte,0o644)
}
