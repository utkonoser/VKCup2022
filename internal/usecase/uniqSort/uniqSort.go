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

// SetVar - множество уникальных значений
var SetVar uniq.Set

// RunUniqSort - основная функция кейса с сортировкой уникальных значений
func RunUniqSort(path string) {
	RunReadAllUniqSort(path)
	items := SetVar.Items()
	sort.QSort(items)
	sort.CreateTxtWithQuickSort(items, path)
}

// RunReadAllUniqSort - функция читает значения в нужной папке и добавляет их в уникальное множество
func RunReadAllUniqSort(path string) {
	if _, err := os.Stat(path + "res.txt"); err == nil {
		err = os.Remove(path + "res.txt")
		if err != nil {
			log.Fatal(err)
		}
	}

	var wg sync.WaitGroup

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, txtFile := range files {
		txtFile := txtFile

		wg.Add(1)
		go func() {
			file, err := os.Open(path + txtFile.Name())
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
