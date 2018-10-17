package main

import (
	"log"
	"os"
	"sync"
)

// thumbnail comes from "gopl.io/tree/master/ch8/thumbnail"

// makeThumbnailsSynchronously original makes thumbnails of specified files - it does this synchronously
// so we're going to take it and use concurrency to process each file
func makeThumbnailsSynchronously(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

// makeThumbnailsConcurrently makes thumbnails of the specified files in parallel
func makeThumbnailsConcurrently(filenames []string) {
	ch := make(chan struct{})

	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f) // NOTE: ignoring errors
			ch <- struct{}{}
		}(f)
	}

	// wait for go routines to complete
	for range filenames {
		<-ch
	}
}

// makeThumbnailsConcurrently2 makes thumbnails for each file received from the channel.
// It returns the number of bytes occupied by the files it creates.
// Here we also use waitgroups.
func makeThumbnailsConcurrently2(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of working goroutines

	for f := range filenames {
		wg.Add(1)
		// worker
		go func(f) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // okay to ignore error
			sizes <- info.Size()
		}(f)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += sizes
	}

	return total
}
