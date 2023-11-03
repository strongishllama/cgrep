package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		log.Fatalln("invalid arguments, expected cgrep <expression>")
	}
	file, err := os.Stdin.Stat()
	if err != nil {
		log.Fatalf("failed to stat stdin: %v\n", err)
	}
	if file.Size() == 0 {
		log.Fatalln("stdin is empty, nothing to parse")
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		if strings.Contains(text, os.Args[1]) {
			text = strings.ReplaceAll(text, os.Args[1], fmt.Sprintf("\033[1;31m%s\033[0m", os.Args[1]))
		}

		fmt.Println(text)
	}

	if scanner.Err() != nil {
		log.Fatalln(scanner.Err())
	}
}
