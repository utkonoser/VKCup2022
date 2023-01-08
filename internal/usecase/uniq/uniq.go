package uniq

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

// SetInterface - интерфейс уникального множества
type SetInterface interface {
	Insert(item int)
	In(item int) bool
	Items() []int
}

var SetVar SetStruct

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
func (s *SetStruct) Items() []int {
	var items []int
	for item := range s.items {
		items = append(items, item)
	}
	return items
}

// RunReadAllUniq - функция, читает файлы из нужной папки и сразу записывает уникальные значения в результирующий файл
func RunReadAllUniq(path string) {
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
				SetVar.Lock()
				SetVar.Insert(num)
				SetVar.Unlock()
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}()
	}
	wg.Wait()
}

func RunUniq(path string) {

	RunReadAllUniq(path)
	items := SetVar.Items()
	CreateTxtWithUniq(items, path)
}

func CreateTxtWithUniq(items []int, path string) {

	resTxt, err := os.Create(path + "res.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = resTxt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	for _, i := range items {
		_, err = resTxt.WriteString(fmt.Sprintln(i))
		if err != nil {
			log.Fatal(err)
		}
	}
}
