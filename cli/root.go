package cli

import (
	"fmt"
	"godownload/internal"
	"sync"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

// Services
var (
	wg sync.WaitGroup
)

// Flags
var (
	url       string
	output    string
	numChunks int
)

// CLI Commands
var (
	rootCmd = &cobra.Command{
		Use:   "godownload",
		Short: "Concurrent downloader",
		Long:  "Godownload is a local concurrent downloader",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Godownload -- Concurrent downloader v0.1.0")
		},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Godownload",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Godownload v0.1.0")
		},
	}

	downloadCmd = &cobra.Command{
		Use:   "download",
		Short: "Download the file in the given link",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Program started!")
			size, _ := internal.GetFileMetadata(url)
			chunks := internal.CreateChunks(size, numChunks)
			bar := progressbar.DefaultBytes(
				size,
				"downloading",
			)
			for _, c := range chunks {
				wg.Add(1)
				go internal.DownloadChunk(url, c, bar, &wg)
			}
			wg.Wait()
			fmt.Println("Downloaded successfully!")
			internal.MergeChunks(chunks, output)
			internal.CleanUpTemp(chunks)
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	downloadCmd.PersistentFlags().StringVarP(&url, "url", "u", "", "the url you want to download")
	downloadCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "output file name")
	downloadCmd.PersistentFlags().IntVarP(&numChunks, "chunk", "c", 20, "the number fo goroutines")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(downloadCmd)
}
