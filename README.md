# ASCII Art Banner Generator

This Go program generates ASCII art banners from text using a specified font. The output is written to a file specified by the `--output` flag.

## Objectives

- The output must be written into a file named by using the `--output=<fileName.txt>` flag.
- Any incorrect usage should return the usage message:
```bash
Usage: go run . --output=<fileName.txt> [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard
```
- The program should handle ASCII art conversions correctly and write to the specified output file.

## Prerequisites

- Go 1.16 or higher

## Project Structure

ascii_output/   
├── main.go  
├── output/   
│ └── output.go   
└── README.md     


## Usage

### Running the Program

To run the program, use the following command:

```sh
go run . --output=<fileName.txt> <message> <bannerFile>
```
- `--output=<fileName.txt>:` Specifies the output file where the ASCII art will be written.
- `<message>:` The message to be converted into ASCII art.
- `<bannerFile>:` The file containing the ASCII banner font.

## Examples
1. Generate an ASCII banner for "hello" using the "standard" font and write it to banner.txt:
```bash
go run . --output=banner.txt "hello" standard
```
2. View the output:
```bash
cat banner.txt
```
## Incorrect Usage
If the --output flag is missing or incorrectly formatted, or if not enough arguments are provided, the program will display the usage message:

```sh
Usage: go run . --output=<fileName.txt> [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard
```
## Testing
It is recommended to create unit tests to ensure the functions in the output package work correctly. Use Go's testing framework for this purpose.

## License
This project is open-source and available under the MIT License.







