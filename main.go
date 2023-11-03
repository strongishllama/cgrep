package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatalln("invalid arguments, expected expression to search for")
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
