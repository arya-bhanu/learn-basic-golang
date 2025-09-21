package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var (
	wg sync.WaitGroup
)

type FileInfo struct {
	path string
	sum  []byte
	hex  string
}

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

func proceedAsync() {
	readChan := make(chan FileInfo, 10)
	go func() {
		err := filepath.Walk(tempPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err.Error())
				return err
			}

			if info.IsDir() {
				return nil
			}

			fileOpened, err := os.ReadFile(path)

			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			readChan <- FileInfo{
				sum:  fileOpened,
				path: path,
			}

			return nil
		})

		if err != nil {
			fmt.Print(err.Error())
		}

		close(readChan)
	}()

	// antrian aliran 1
	mdW1 := getMD5Sum(readChan)
	// antiran aliran 2
	mdW2 := getMD5Sum(readChan)
	// antrian aliran 3
	mdW3 := getMD5Sum(readChan)

	// gabung semua aliran ke chanel yang lebih besar dan global sehingga bisa di consume secara fleksibel, tidak terbatas pada 1 aliran saja per consumer
	mergedChannelMD5Sum := mergeMD5SumChanel(mdW1, mdW2, mdW3)

	var wg2 sync.WaitGroup

	renameFile(mergedChannelMD5Sum, &wg2)
	renameFile(mergedChannelMD5Sum, &wg2)
	renameFile(mergedChannelMD5Sum, &wg2)

	wg2.Wait()
}

func renameFile(chanIn <-chan FileInfo, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for fileInfo := range chanIn {
			newPath := filepath.Join(tempPath, fmt.Sprintf("file-%s.txt", fileInfo.hex))
			err := os.Rename(fileInfo.path, newPath)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}()
}

func getMD5Sum(chanIn <-chan FileInfo) <-chan FileInfo {
	chanString := make(chan FileInfo, 10)

	go func() {
		defer close(chanString)
		for valBytes := range chanIn {
			md5Sum := md5.Sum(valBytes.sum)
			valBytes.hex = hex.EncodeToString(md5Sum[:])
			chanString <- valBytes
		}
	}()

	return chanString
}

func mergeMD5SumChanel(currents ...<-chan FileInfo) <-chan FileInfo {
	mainRiverChanel := make(chan FileInfo, 10)
	for _, channel := range currents {
		wg.Add(1)
		go func(x <-chan FileInfo) {
			defer wg.Done()
			for value := range x {
				mainRiverChanel <- value
			}
		}(channel)
	}

	// fungsi pengawaan dan watching go routine yang sedang berjalan dilakukan secara paralel sehingga tidak memblokir aliran mainRiverChanel
	go func() {
		wg.Wait()
		close(mainRiverChanel)
	}()

	return mainRiverChanel
}
