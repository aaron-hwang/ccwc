package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/aaron-hwang/cwcc/src/word_counter"
	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("wc", "Outputs the word count and other metadata of the given input.")
	shouldCountBytes := parser.Flag("c", "countBytes", &argparse.Options{Required: false})
	file := parser.FilePositional(os.O_RDONLY, 0700, &argparse.Options{Default: nil})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(parser.Usage(err))
		return
	}
	if *shouldCountBytes {
		fmt.Println("test")
	}

	reader := bufio.NewReader(os.Stdin)
	if file != nil {
		reader = bufio.NewReader(file)
	}
	wc := word_counter.NewWordCounter(reader)
	wc.ReadBytes()
	println(wc.ByteCount)
}
