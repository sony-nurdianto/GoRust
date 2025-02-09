package main

import (
	"log"
	"os"
	"sync"

	"github.com/sony-nurdianto/GoRust/thumbnails/utils"
)

func makeThumbnails(filename []string) {
	for _, f := range filename {
		if _, err := utils.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

func makeThumbnails2(filesname []string) {
	for _, f := range filesname {
		go utils.ImageFile(f)
	}
}

func makeThumbnails3(filename []string) {
	ch := make(chan struct{})
	for _, f := range filename {
		go func(f string) {
			utils.ImageFile(f)
			ch <- struct{}{}
		}(f)
	}
	for range filename {
		<-ch
	}
}

func makeThumbnails4(filesnames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbnails string
		err        error
	}
	ch := make(chan item, len(filesnames))
	for _, f := range filesnames {
		go func(f string) {
			var it item
			it.thumbnails, it.err = utils.ImageFile(f)
			ch <- it
		}(f)
	}

	for range filesnames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbnails)

	}
	return thumbfiles, nil
}

func makeThumbnails5(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumb, err := utils.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}

func main() {
}
