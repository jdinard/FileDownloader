package main

import (
	"downloader"
	"flag"
	"fmt"
)

func main() {
	// Use the golang flag package to get command line arguments, set defaults, and support standard things like usage/help
	remoteURL := *flag.String("source_url:", "", "a fully qualified URL of the file to download")
	outfile := *flag.String("output_file:", "file.out", "The name of the file to output to")
	numberOfChunks := *flag.Int("num_chunks:", 4, "The number of file chunks to download")
	chunkSize := *flag.Int("chunk_size:", 1048576, "The size of the chunks to download")
	flag.Parse()

	//Has to at least have http://a.a
	if len(remoteURL) < 9 {
		fmt.Println("You must specify the source Url with the argument -source_url=<url>")
	} else {
		// Download the file in chunks
		downloader.DownloadFileInChunks(numberOfChunks, chunkSize, remoteURL, outfile)
	}
}
