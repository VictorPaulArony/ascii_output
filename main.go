package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"ascii_output/output"
)

func main() {
	// Define the flag
	outputFile := flag.String("output", "", "output banner file")
	flag.Parse()

	// Check if the output flag is provided correctly
	if *outputFile == "" {
		fmt.Println("Usage: go run . --output=<fileName.txt> [STRING] [BANNER]")
		fmt.Println("EX: go run . --output=<fileName.txt> something standard")
		return
	}

	// Ensure there are enough arguments after parsing flags
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Usage: go run . --output=<fileName.txt> [STRING] [BANNER]")
		fmt.Println("EX: go run . --output=<fileName.txt> something standard")
		return
	}

	// Capture the message and banner file arguments
	message := args[0]
	bannerFile := args[1]

	// Read the banner file
	data, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Println("Error reading banner file:", err)
		return
	}
	lines := strings.Split(string(data), "\n")

	// Display the message in ASCII art and write to the specified output file
	output.Display(message, lines, *outputFile)
}
