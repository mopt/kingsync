package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"github.com/mopt/kingsync/syncers"
	"log"
	"os"
	"strings"
)

func main() {
	filenames := flag.String("filenames", "", "file containing filenames to sync (omit for stdin)")
	target := flag.String("target", "", "target sync location")
	flag.Parse()

	if *target == "" {
		log.Fatal("Missing target sync location")
	}

	// open up the input source
	var reader *bufio.Reader
	if *filenames != "" {
		inputFiles, err := os.Open(*filenames)
		if err != nil {
			log.Fatal("Couldn't open input filenames:", err)
		}
		defer inputFiles.Close()
		reader = bufio.NewReader(inputFiles)
	} else {
		reader = bufio.NewReader(os.Stdin)
		fmt.Println("Paste in your list of files to sync:")
	}

	// read in all of the input files
	var err error
	files := make([]string, 0, 64)
	for err == nil {
		var line string
		line, err = reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "" {
			break
		}

		files = append(files, line)
	}

	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	fmt.Printf("Received %d files.\n", len(files))

	syncer := syncers.SyncerNop{}
	progress := make(chan int)
	syncer.Sync(files, *target, progress)

	for {
		percent := <-progress
		fmt.Printf("%d%%\n", percent)

		if percent >= 100 {
			break
		}

		if percent < 0 {
			log.Fatal("TODO: syncer died")
		}
	}

}
