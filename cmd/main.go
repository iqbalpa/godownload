package main

import (
	"fmt"
	"godownload/internal"
	"sync"
)

func main() {
	fmt.Println("helo")

	var wg sync.WaitGroup

	url := "https://www.sample-videos.com/video321/mp4/720/big_buck_bunny_720p_30mb.mp4"
	size, _ := internal.GetFileMetadata(url)
	chunks := internal.CreateChunks(size, 20)

	for _, c := range chunks {
		wg.Add(1)
		go internal.DownloadChunk(url, c, &wg)
	}

	wg.Wait()
	fmt.Println("Downloaded successfully!")

	internal.MergeChunks(chunks, "result.mp4")
}
