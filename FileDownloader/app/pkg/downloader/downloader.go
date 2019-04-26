package downloader

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var writeBuffer []string
var group sync.WaitGroup

// DownloadFileInChunks -- (numberOfChunks, sizeOfChunks, remoteFile, outFile) Downloads the first n chunks of a file in parallel and then concatenates it
func DownloadFileInChunks(numberOfChunks int, sizeOfChunks int, remoteFile string, outfile string) {
	writeBuffer = make([]string, numberOfChunks)

	// Create a bunch of goroutines that concurrently download each chunk of the file
	for i := 0; i < numberOfChunks; i++ {
		group.Add(1)
		start, end := calculateByteRange(i, sizeOfChunks)
		go downloadFileChunk(start, end, remoteFile, outfile, i)
	}
	// Wait for all routines to finish
	group.Wait()

	// Concatenate the results of the files together
	concatenateTempFiles(numberOfChunks, outfile)
}

func calculateByteRange(iteration int, sizeOfChunks int) (int, int) {
	return calculateByteStart(iteration, sizeOfChunks), calculateByteEnd(iteration, sizeOfChunks)
}

func calculateByteStart(iteration int, sizeOfChunks int) int {
	return iteration * sizeOfChunks
}

func calculateByteEnd(iteration int, sizeOfChunks int) int {
	return ((iteration+1)*sizeOfChunks - 1)
}

func getTempFileName(chunk int, outfile string) string {
	return outfile + "-" + strconv.Itoa(chunk)
}

func concatenateTempFiles(numberOfChunks int, outfile string) {
	// Delete the file if it exists before opening it/creating it
	os.Remove(outfile)
	output, err := os.OpenFile(outfile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Fatalln("Failed to open output file ", err)
	}

	// Concatenate the files into one
	for x := 0; x < numberOfChunks; x++ {
		tmpFileName := getTempFileName(x, outfile)

		tempFile, errIn := os.Open(tmpFileName)

		if errIn != nil {
			log.Fatalln("Failed to open temporary file for reading:", err)
		}

		_, errCopy := io.Copy(output, tempFile)
		if errCopy != nil {
			log.Fatalln("Failed to append temporary file: ", err)
		}

		defer tempFile.Close()
		os.Remove(tmpFileName)
	}

	fmt.Println("File saved as " + outfile)

	defer output.Close()
}

func downloadFileChunk(startByte int, endByte int, remoteFile string, outfile string, fileChunk int) {
	// Setup a client for the remote file
	client := &http.Client{}
	req, _ := http.NewRequest("GET", remoteFile, nil)

	// Calculate the range of bytes to be downloaded for the file and create a range header
	rangeHeader := "bytes=" + strconv.Itoa(startByte) + "-" + strconv.Itoa(endByte)
	fmt.Println("Downloading " + rangeHeader)
	req.Header.Add("Range", rangeHeader)

	// Get the range of bytes
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln("Could not process request to download file chunk", err)
	}
	defer resp.Body.Close()
	reader, _ := ioutil.ReadAll(resp.Body)

	// Buffer and write the chunk to a temporary file
	writeBuffer[fileChunk] = string(reader)
	ioutil.WriteFile(getTempFileName(fileChunk, outfile), []byte(string(writeBuffer[fileChunk])), 0x777)

	// Let the waitgroup know this routine is done
	group.Done()
}
