package uniqSort

import (
	"bufio"
	"goElimination/internal/usecase/sort"
	"goElimination/internal/usecase/uniq"
	"log"
	"os"
	"strconv"
	"sync"
)

var SetVar uniq.Set

func RunUniqSort(path string) {
	RunReadAllUniqSort()
	items := SetVar.Items()
	sort.QSort(items)
	sort.CreateTxtWithQuickSort(items, path)
}

func RunReadAllUniqSort() {
	if _, err := os.Stat("../../data/res.txt"); err == nil {
		err = os.Remove("../../data/res.txt")
		if err != nil {
			log.Fatal(err)
		}
	}

	var wg sync.WaitGroup

	files, err := os.ReadDir("../../data/")
	if err != nil {
		log.Fatal(err)
	}
	for _, txtFile := range files {
		txtFile := txtFile

		wg.Add(1)
		go func() {
			file, err := os.Open("../../data/" + txtFile.Name())
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
						SetVar.Lock()
						if ok := SetVar.In(num); !ok {
							SetVar.Insert(num)
						}
						SetVar.Unlock()
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
