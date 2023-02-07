package uniq

import (
	"bufio"
	"fmt"
	"goElimination/internal/usecase/sort"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

// RunUniq - функция, запускающая кейс с уникальным множеством
func RunUniq(path string) {
	var SetVar SetStruct

	fmt.Println("Run case 'uniq'...")
	startCaseUniq := time.Now()
	SetVar.RunReadAllUniq(path)
	endRunReadAllFunc := time.Since(startCaseUniq)
	fmt.Println("End of reading all files and insert in set: ", endRunReadAllFunc)

	items := SetVar.Items().ShowItems()

	startCreateTxt := time.Now()
	sort.CreateTxt(items, path)
	endCreateTxt := time.Since(startCreateTxt)
	fmt.Println("End of creating res.txt: ", endCreateTxt)

	endCaseUniq := time.Since(startCaseUniq)
	fmt.Println("Elapsed time for case 'uniq': ", endCaseUniq)
}

// SetInterface - интерфейс уникального множества
type SetInterface interface {
	Insert(item int)
	In(item int) bool
	Items() sort.SliceItems
}

// SetStruct - реализация сета уникальных значений через map
type SetStruct struct {
	items map[int]struct{}
	sync.Mutex
}

// Insert - метод SetStruct, который добавляет элемент в уникальное множество
func (s *SetStruct) Insert(item int) {
	if s.items == nil {
		s.items = make(map[int]struct{})
	}
	_, ok := s.items[item]
	if !ok {
		s.items[item] = struct{}{}
	}
}

// In - метод SetStruct, проверяющий находится ли элемент во множестве
func (s *SetStruct) In(item int) bool {
	_, ok := s.items[item]
	return ok
}

// Items - метод SetStruct, возвращающий слайс элементов из уникального множества
func (s *SetStruct) Items() sort.SliceItems {
	var items sort.Result
	for item := range s.items {
		items.Res = append(items.Res, item)
	}
	return &items
}

// RunReadAllUniq - функция, читает файлы из нужной папки и добавляет числа из файлов в уникальное множество
func (s *SetStruct) RunReadAllUniq(path string) {
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

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				num, err := strconv.Atoi(scanner.Text())
				if err != nil {
					log.Fatal(err)
				}
				s.Lock()
				s.Insert(num)
				s.Unlock()
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}()
	}
	wg.Wait()
}
