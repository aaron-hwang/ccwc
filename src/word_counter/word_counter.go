package word_counter

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type WordCounter struct {
	ByteCount int
	LineCount int
	WordCount int
	curr      *os.File
	reader    *bufio.Reader
	scanner   *bufio.Scanner
}

func NewWordCounter(file *os.File) *WordCounter {
	// Go will init the non defined values to their nil (0) value.
	return &WordCounter{reader: bufio.NewReader(file), scanner: bufio.NewScanner(file), curr: file}
}

func (wc *WordCounter) reInitCounter() {
	wc.curr.Seek(0, io.SeekStart)
	wc.reader = bufio.NewReader(wc.curr)
	wc.scanner = bufio.NewScanner(wc.curr)
}

// Counts the number of bytes in the counter's current reader.
func (wc *WordCounter) ReadBytes() (int, error) {
	wc.reInitCounter()
	res := 0
	for {
		temp := make([]byte, 1)
		n, err := wc.reader.Read(temp)
		res += n
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Could not read file")
			return 0, fmt.Errorf("could not read file")
		}
	}
	wc.ByteCount = res
	return wc.ByteCount, nil
}

func (wc *WordCounter) ReadLineCount() (int, error) {
	wc.reInitCounter()
	lines := 0
	for {
		_, _, err := wc.reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error reading lines of file.")
			return 0, fmt.Errorf("error reading lines of file")
		}
		lines += 1
	}
	wc.LineCount = lines
	return lines, nil
}

func (wc *WordCounter) ReadWordCount() (int, error) {
	wc.reInitCounter()
	words := 0
	wc.scanner.Split(bufio.ScanWords)
	for wc.scanner.Scan() {
		words += 1
	}

	wc.WordCount = words
	return words, nil
}

// TODO: Change to do in one pass.
func (wc *WordCounter) ReadAll() {
	wc.ReadBytes()
	wc.ReadLineCount()
	wc.ReadWordCount()
}
