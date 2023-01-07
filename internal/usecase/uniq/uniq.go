package uniq

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

var SetVar Set

type Set struct {
	items map[int]struct{}
	sync.Mutex
}

func (s *Set) Insert(item int) {
	if s.items == nil {
		s.items = make(map[int]struct{})
	}
	_, ok := s.items[item]
	if !ok {
		s.items[item] = struct{}{}
	}
}

func (s *Set) In(item int) bool {
	_, ok := s.items[item]
	return ok
}

func (s *Set) Items() []int {
	var items []int
	for item := range s.items {
		items = append(items, item)
	}
	return items
}

func RunUniq() {
	if _, err := os.Stat("../../data/res.txt"); err == nil {
		err = os.Remove("../../data/res.txt")
		if err != nil {
			log.Fatal(err)
		}
	}

	resTxt, err := os.Create("../../data/res.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = resTxt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

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
							_, err = resTxt.WriteString(fmt.Sprintln(num))
							if err != nil {
								log.Fatal(err)
							}
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
