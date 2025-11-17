package main

import (
	"fmt"
	"os"

	"github.com/aaron-hwang/cwcc/src/word_counter"
	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("wc", "Outputs the word count and other metadata of the given input.")
	shouldCountBytes := parser.Flag("c", "bytes", &argparse.Options{Required: false})
	shouldCountLines := parser.Flag("l", "lines", &argparse.Options{Required: false})
	shouldCountWords := parser.Flag("w", "words", &argparse.Options{Required: false})
	file := parser.FilePositional(os.O_RDONLY, 0700, &argparse.Options{Default: nil})

	// Always call parser.Parse before using parser flags.
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(parser.Usage(err))
		return
	}

	toRead := os.Stdin
	if file != nil {
		toRead = file
	}

	wc := word_counter.NewWordCounter(toRead)
	if *shouldCountBytes {
		wc.ReadBytes()
		println(wc.ByteCount)
	}
	if *shouldCountLines {
		wc.ReadLineCount()
		println(wc.LineCount)
	}
	if *shouldCountWords {
		wc.ReadWordCount()
		println(wc.WordCount)
	}
	if !*shouldCountBytes && !*shouldCountLines && !*shouldCountWords {
		wc.ReadAll()
		fmt.Printf("%d %d %d %s", wc.ByteCount, wc.LineCount, wc.WordCount, toRead.Name())
	}
}
