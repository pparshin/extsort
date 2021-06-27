package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/pparshin/extsort/internal/ioutil"
	"github.com/pparshin/extsort/internal/merge"
	"github.com/pparshin/extsort/internal/sort"
)

var (
	input  = flag.String("in", "input.txt", "Path to the input file")
	output = flag.String("out", "output.txt", "Path to the output file")
)

const (
	chunkSize      = 100
	mergeThreshold = 10
)

func main() {
	flag.Parse()

	printMemUsage()

	in, err := ioutil.ReadLineByLine(*input)
	if err != nil {
		log.Fatal(err)
	}

	merger := merge.NewMerger()
	chunk := sort.NewChunk(chunkSize)
	for line := range in {
		chunk.Add(line)

		if chunk.Len() == chunkSize {
			var filename string
			filename, err = ioutil.SaveChunkToTempFile(chunk)
			if err != nil {
				log.Fatal(err)
			}
			merger.Add(filename)

			chunk = sort.NewChunk(chunkSize)

			printMemUsage()
		}

		if merger.Len() == mergeThreshold {
			err = merger.Merge()
			if err != nil {
				log.Fatal(err)
			}

			printMemUsage()
		}
	}

	if chunk.Len() != 0 {
		var filename string
		filename, err = ioutil.SaveChunkToTempFile(chunk)
		if err != nil {
			log.Fatal(err)
		}
		merger.Add(filename)
	}

	result, err := merger.MergeAll()
	if err != nil {
		log.Fatal(err)
	}

	err = os.Rename(result, *output)
	if err != nil {
		log.Fatal(err)
	}

	printMemUsage()
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
