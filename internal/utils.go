package internal

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func GetFileMetadata(url string) (int64, error) {
	resp, err := http.Head(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("server error: %v", resp.Status)
	}

	if resp.Header.Get("Accept-Ranges") != "bytes" {
		return 0, fmt.Errorf("server does not support partial content downloads")
	}

	size, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid content-length: %v", err)
	}

	return size, nil
}

func CleanUpTemp(chunks []Chunk) error {
	fmt.Println("Cleaning up temp files...")
	for _, c := range chunks {
		err := os.Remove(c.Filename)
		if err != nil {
			fmt.Printf("Failed to delete %s\n", c.Filename)
			return err
		}
	}
	fmt.Println("Cleaned up all the temporary files")
	return nil
}
