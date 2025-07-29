# GoDownload

GoDownload is a simple and efficient concurrent file downloader written in Go. It splits a file into multiple chunks and downloads them in parallel, which can significantly speed up the download process for large files. After downloading, it merges the chunks back into a single file.

## How it Works

1.  **Get File Metadata**: It first retrieves the file size from the given URL.
2.  **Create Chunks**: The file is then divided into a predefined number of chunks.
3.  **Concurrent Download**: Each chunk is downloaded concurrently using goroutines.
4.  **Merge Chunks**: Once all chunks are downloaded, they are merged in the correct order to reconstruct the original file.

## Getting Started

### Prerequisites

- Go (version 1.24.4 or higher)

### Running the Project

1.  Clone the repository:
    ```bash
    git clone <repository-url>
    cd godownload
    ```

2.  Run the application:
    ```bash
    go run cmd/main.go
    ```

    This will download a sample video file and save it as `result.mp4` in the project's root directory.

## Project Structure

-   `cmd/main.go`: The main entry point of the application.
-   `internal/`: Contains the core logic of the downloader.
    -   `downloader.go`: Handles the downloading of individual chunks.
    -   `chunk.go`: Manages the creation and merging of file chunks.
    -   `utils.go`: Provides utility functions, such as retrieving file metadata.
-   `go.mod`: Defines the project's module and dependencies.
