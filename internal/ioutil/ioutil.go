package ioutil

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"

	"github.com/pparshin/extsort/internal/sort"
)

func CreateTempFile() (*os.File, error) {
	return ioutil.TempFile("", "chunk")
}

func SaveChunkToTempFile(chunk *sort.Chunk) (filename string, err error) {
	f, err := CreateTempFile()
	if err != nil {
		return "", err
	}
	defer f.Close()
	defer func() {
		if err != nil {
			cleanupErr := os.Remove(f.Name())
			if cleanupErr != nil {
				log.Println(cleanupErr)
			}
		}
	}()

	w := bufio.NewWriter(f)
	content := chunk.ToArray()
	for _, line := range content {
		_, err = w.WriteString(line + "\n")
		if err != nil {
			return "", err
		}
	}

	err = w.Flush()
	if err != nil {
		return "", err
	}

	return f.Name(), nil
}

func ReadLineByLine(filename string) (chan string, error) {
	source, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	ch := make(chan string)

	go func() {
		scanner := bufio.NewScanner(source)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		close(ch)

		source.Close()
	}()

	return ch, nil
}
