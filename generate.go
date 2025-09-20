package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
)

const totalFile = 50000
const contentLength = 100000

var (
	tempPath string
)

func init() {
	rand.Seed(time.Now().UnixNano())
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	tempPath = filepath.Join(os.Getenv("TEMP"), "pipeline-generated")
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
	for t := range totalFile {
		filename := path.Join(tempPath, fmt.Sprintf("file-%d-sample.txt", t))
		content := randomString(contentLength)
		err := os.WriteFile(filename, []byte(content), os.ModePerm)
		if err != nil {
			fmt.Println("Error writing file", filename)
		}
		if t%100 == 0 && t > 0 {
			fmt.Printf("Current files created: %d\n", t)
		}
	}
	fmt.Printf("Total files created: %d\n", totalFile)
}

func randomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
