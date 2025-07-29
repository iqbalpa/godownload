# GoDownload

GoDownload is a simple and efficient concurrent file downloader written in Go. It splits a file into multiple chunks and downloads them in parallel, which can significantly speed up the download process for large files. After downloading, it merges the chunks back into a single file.

This tool provides a command-line interface (CLI) to specify the download URL, the output file name, and the number of concurrent connections. It also features a progress bar to monitor the download progress.

## How it Works

1.  **Parse CLI Flags**: It parses the command-line arguments for the URL, output file name, and number of connections.
2.  **Get File Metadata**: It first retrieves the file size from the given URL.
3.  **Create Chunks**: The file is then divided into a predefined number of chunks.
4.  **Concurrent Download**: Each chunk is downloaded concurrently using goroutines, with a progress bar showing the overall progress.
5.  **Merge Chunks**: Once all chunks are downloaded, they are merged in the correct order to reconstruct the original file.

## Getting Started

### Prerequisites

- Go (version 1.24.4 or higher)

### Running the Project

1.  Clone the repository:
    ```bash
    git clone <repository-url>
    cd godownload
    ```

2.  Build the application:
    ```bash
    go build -o godownload cmd/main.go
    ```

3.  Run the application with the desired flags:
    ```bash
    ./godownload -url <file-url> -o <output-filename> -n <number-of-connections>
    ```
    For example:
    ```bash
    ./godownload -url "https://sample-videos.com/video123/mp4/720/big_buck_bunny_720p_1mb.mp4" -o result.mp4 -n 10
    ```

## Project Structure

-   `cmd/main.go`: The main entry point of the application.
-   `cli/root.go`: Defines the CLI commands and flags using the Cobra library.
-   `internal/`: Contains the core logic of the downloader.
    -   `downloader.go`: Handles the downloading of individual chunks.
    -   `chunk.go`: Manages the creation and merging of file chunks.
    -   `utils.go`: Provides utility functions, such as retrieving file metadata.
-   `go.mod`: Defines the project's module and dependencies.
