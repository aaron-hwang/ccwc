package word_counter

import (
	"bufio"
	"io"
	"log"
)

type WordCounter struct {
	ByteCount int
	LineCount int
	WordCount int
	reader    bufio.Reader
}

func NewWordCounter(reader io.Reader) *WordCounter {
	// Go will init the non defined values to their nil (0) value.
	return &WordCounter{reader: *bufio.NewReader(reader)}
}

func (wc *WordCounter) ReadBytes() int {
	for {
		temp := make([]byte, 5)
		n, err := wc.reader.Read(temp)
		wc.ByteCount += n
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Could not read file")
		}
	}
	return wc.ByteCount
}

func ReadLines() {

}

func ReadWordCount() {

}
