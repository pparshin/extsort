package main

import (
	"bufio"
	"flag"
	"log"
	"math/rand"
	"os"
	"path"
	"time"
)

var (
	maxLen = flag.Uint("max_len", 10, "Maximum length of the string")
	lines  = flag.Uint("lines", 1000, "Total number of strings to generate")
	out    = flag.String("out", "input.txt", "The output file name")
)

var (
	alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
)

func main() {
	flag.Parse()

	f, err := os.Create(getOutputPath())
	if err != nil {
		log.Fatal(err)
	}

	w := bufio.NewWriter(f)
	for line := range generate() {
		_, err = w.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	err = w.Flush()
	if err != nil {
		log.Fatal(err)
	}
}

func getOutputPath() string {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return path.Join(workingDir, *out)
}

func generate() <-chan string {
	rand.Seed(time.Now().UnixNano())

	ch := make(chan string)

	go func() {
		for i := uint(0); i < *lines; i++ {
			ch <- generateRandomString()
		}
		close(ch)
	}()

	return ch
}

func generateRandomString() string {
	length := rand.Intn(int(*maxLen)) + 1
	b := make([]rune, length)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(b)
}
