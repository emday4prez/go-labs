package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	linesFlag := flag.Bool("l", false, "Count lines")
	wordsFlag := flag.Bool("w", false, "Count words")
	bytesFlag := flag.Bool("c", false, "Count bytes")
	flag.Parse()

	var reader io.Reader

	if len(flag.Args()) > 0 {
		filename := flag.Args()[0]
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "gocount: error opening file: %v\n", err)
			os.Exit(1)
		}
		// IMPORTANT: Ensure the file is closed when main exits.
		defer file.Close()
		reader = file
	} else {
		reader = os.Stdin
	}

	lines, words, bytes := count(reader)

	noFlagsSet := !*linesFlag && !*wordsFlag && !*bytesFlag

	if *bytesFlag || noFlagsSet {
		fmt.Printf("Bytes: %d\n", bytes)
	}
	if *wordsFlag || noFlagsSet {
		fmt.Printf("Words: %d\n", words)
	}
	if *linesFlag || noFlagsSet {
		fmt.Printf("Lines: %d\n", lines)
	}
}

func count(r io.Reader) (int, int, int) {
	scanner := bufio.NewScanner(r)

	lines := 0
	words := 0
	byteCount := 0

	for scanner.Scan() {
		lines++

		lineBytes := scanner.Bytes()
		byteCount += len(lineBytes)
		words += len(bytes.Fields(lineBytes))

	}

	byteCount += lines
	return lines, words, byteCount

}
