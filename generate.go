package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

const totalFile = 50000
const contentLength = 100000

var (
	tempPath string
	mu       sync.Mutex
)

func init() {
	rand.Seed(time.Now().UnixNano())
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	tempPath = filepath.Join(os.Getenv("TEMP"), "pipeline-generated")
}

type GeneratedFile struct {
	content  string
	filename string
}

func generateFiles() {

	err := os.RemoveAll(tempPath)
	if err != nil {
		fmt.Printf("Error remove all path: %s", err.Error())
	}
	err = os.MkdirAll(tempPath, os.ModePerm)
	if err != nil {
		fmt.Printf("Error mkdir related path: %s", err.Error())
	}
	filenameChan := make(chan GeneratedFile, 3)
	totalFileCreatedChan := 0
	go func() {
		defer close(filenameChan)
		for t := range totalFile {
			filenameChan <- GeneratedFile{
				filename: path.Join(tempPath, fmt.Sprintf("file-%d-sample.txt", t)),
			}
		}
	}()

	var wg1 sync.WaitGroup
	// worker generate file
	generatedChan1 := generateFileWorker(contentLength, filenameChan)
	generatedChan2 := generateFileWorker(contentLength, filenameChan)
	generatedChan3 := generateFileWorker(contentLength, filenameChan)
	generatedChan4 := generateFileWorker(contentLength, filenameChan)
	generatedChan5 := generateFileWorker(contentLength, filenameChan)

	chanFiles := waitAllGenerateFileWorker(&wg1, generatedChan1, generatedChan2, generatedChan3, generatedChan4, generatedChan5)

	var wg sync.WaitGroup
	// worker write file
	writeFileWorker(chanFiles, &totalFileCreatedChan, &wg)
	writeFileWorker(chanFiles, &totalFileCreatedChan, &wg)
	writeFileWorker(chanFiles, &totalFileCreatedChan, &wg)
	writeFileWorker(chanFiles, &totalFileCreatedChan, &wg)
	writeFileWorker(chanFiles, &totalFileCreatedChan, &wg)
	wg.Wait()
	fmt.Printf("Total files created: %d\n", totalFileCreatedChan)
}

func generateFileWorker(contentLength int, filename <-chan GeneratedFile) <-chan GeneratedFile {
	out := make(chan GeneratedFile, 3)
	go func() {
		defer close(out)
		for currFile := range filename {
			content := randomString(contentLength)
			currFile.content = content
			out <- currFile
		}
	}()
	return out
}

func waitAllGenerateFileWorker(wg *sync.WaitGroup, workers ...<-chan GeneratedFile) <-chan GeneratedFile {
	contentOut := make(chan GeneratedFile, 5)
	for _, worker := range workers {
		wg.Add(1)
		go func(worker <-chan GeneratedFile) {
			defer wg.Done()
			for work := range worker {
				contentOut <- work
			}
		}(worker)
	}
	go func() {
		wg.Wait()
		close(contentOut)
	}()
	return contentOut
}

func writeFileWorker(inChan <-chan GeneratedFile, totalFileCreated *int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range inChan {
			mu.Lock()
			currVal := *totalFileCreated
			err := os.WriteFile(val.filename, []byte(val.content), os.ModePerm)
			currVal++
			*totalFileCreated = currVal
			if err != nil {
				fmt.Println("Error writing file", val.filename)
			}
			mu.Unlock()
		}
	}()
}

func randomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
