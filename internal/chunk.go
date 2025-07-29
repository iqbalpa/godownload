package internal

import "fmt"

type Chunk struct {
	ID        int
	StartByte int64
	EndByte   int64
	Filename  string
}

func CreateChunks(totalSize int64, numChunks int) []Chunk {
	chunkSize := totalSize / int64(numChunks)
	var chunks []Chunk

	for i := 0; i < numChunks; i++ {
		start := int64(i) * chunkSize
		end := start + chunkSize - 1
		if i == numChunks-1 {
			end = totalSize - 1
		}

		chunk := Chunk{
			ID:        i,
			StartByte: start,
			EndByte:   end,
			Filename:  fmt.Sprintf("temp/chunk-%d.tmp", i),
		}
		chunks = append(chunks, chunk)
	}

	return chunks
}
