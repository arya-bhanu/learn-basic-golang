package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

var (
	mu sync.Mutex
)

type FileInfoAsync struct {
	Index       int
	FilePath    string
	WorkerIndex int
	Err         error
}

func init() {
	rand.Seed(time.Now().UnixNano())
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	tempPath = filepath.Join(os.Getenv("TEMP"), "pipeline-generated")
}

func GenerateFileAsync() {
	var wg sync.WaitGroup
	filePathChan := generateFilename()
	results := generateWriteFiles(100, filePathChan, &wg)
	counterTotal := 0
	counterSuccess := 0
	for _, result := range results {
		if result.Err != nil {
			log.Printf("error creating file %s. stack trace: %s", result.FilePath, result.Err)
		} else {
			counterSuccess++
		}
		counterTotal++
	}
	log.Printf("%d/%d of total files created", counterSuccess, counterTotal)
}

func generateFilename() <-chan FileInfoAsync {
	out := make(chan FileInfoAsync)
	go func() {
		defer close(out)
		for i := range totalFile {
			filepath := filepath.Join(tempPath, fmt.Sprintf("file-test-%d.txt", i))
			out <- FileInfoAsync{
				FilePath: filepath,
				Index:    i,
			}
		}
	}()

	return out
}

func generateWriteFiles(workerCount int, producer <-chan FileInfoAsync, wg *sync.WaitGroup) []FileInfoAsync {
	out := []FileInfoAsync{}
	for i := range workerCount {
		wg.Add(1)
		go func(indexWorker int) {
			defer wg.Done()
			for product := range producer {
				content := RandomString(contentLength)
				err := os.WriteFile(product.FilePath, []byte(content), os.ModePerm)
				if err != nil {
					product.Err = err
				}
				mu.Lock()
				out = append(out, product)
				mu.Unlock()
			}
		}(i)
	}

	wg.Wait()

	return out
}
