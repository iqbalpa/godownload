package internal

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

func DownloadChunk(url string, chunk Chunk, wg *sync.WaitGroup) {
	defer wg.Done()

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", chunk.StartByte, chunk.EndByte))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("=== Failed to download", err)
		return
	}
	defer resp.Body.Close()

	f, err := os.Create(chunk.Filename)
	if err != nil {
		fmt.Println("=== Failed to create file", err)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		fmt.Println("=== Failed to write the file", err)
		return
	}
}

func MergeChunks(chunks []Chunk, fpath string) error {
	f, err := os.Create(fpath)
	if err != nil {
		fmt.Println("=== Failed to create file", err)
		return err
	}
	defer f.Close()

	for _, chunk := range chunks {
		data, err := os.Open(chunk.Filename)
		if err != nil {
			fmt.Println("=== Failed to open", err)
			return err
		}

		_, err = io.Copy(f, data)
		if err != nil {
			fmt.Println("=== Failed to append the bytes", err)
			return err
		}
	}

	fmt.Println("Merged files successfully!")
	return nil
}
