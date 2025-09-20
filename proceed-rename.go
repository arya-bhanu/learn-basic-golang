package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
)

func proceedSync() {
	counterTotalFound := 0
	counterTotalRenamed := 0
	// loop the path
	err := filepath.Walk(tempPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// it will stop all
			return err
		}

		if info.IsDir() {
			// it will skip
			return nil
		}

		counterTotalFound++

		fileOpened, err := os.ReadFile(path)

		if err != nil {
			return err
		}

		hashed := md5.Sum(fileOpened)

		stringHash := hex.EncodeToString(hashed[:])

		renamedPath := filepath.Join(tempPath, fmt.Sprintf("file-path-renamed-%s.txt", stringHash))

		err = os.Rename(path, renamedPath)

		if err != nil {
			return err
		}

		counterTotalRenamed++

		return nil

	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Total file founded: %d\n", counterTotalFound)
	fmt.Printf("Total file renamed: %d\n", counterTotalRenamed)
}
