package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"sync"
)

func RunReadAllUniqSort() {
	if _, err := os.Stat("data/res.txt"); err == nil {
		err = os.Remove("data/res.txt")
		if err != nil {
			log.Fatal(err)
		}
	}

	var wg sync.WaitGroup

	files, err := os.ReadDir("data/")
	if err != nil {
		log.Fatal(err)
	}
	for _, txtFile := range files {
		txtFile := txtFile

		wg.Add(1)
		go func() {
			file, err := os.Open("data/" + txtFile.Name())
			if err != nil {
				log.Fatal(err)
			}
			defer func() {
				err = file.Close()
				if err != nil {
					log.Fatal(err)
				}
				wg.Done()
			}()

			if txtFile.Name() != "res.txt" {
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					num, err := strconv.Atoi(scanner.Text())
					if err != nil {
						log.Fatal(err)
					}
					wg.Add(1)
					go func() {
						defer wg.Done()
						set.Lock()
						if ok := set.In(num); !ok {
							set.Insert(num)
						}
						set.Unlock()
					}()
				}
				if err := scanner.Err(); err != nil {
					log.Fatal(err)
				}
			}
		}()
	}
	wg.Wait()
}
