package mock

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
)

// dataPath - путь для сгенерированных файлов
const dataPath = "../../../data/"

// createData - функция для создания тестовых файлов при countFiles=50000
// и countIntegers=1000, размер данных будет ~1Gb
func createData(countFiles, countIntegers int) {
	var wg sync.WaitGroup
	for j := 0; j < countFiles; j++ {
		wg.Add(1)
		j := j
		go func() {
			defer wg.Done()
			path := fmt.Sprintf("%s%v.txt", dataPath, j)

			resTxt, err := os.Create(path)
			if err != nil {
				log.Fatal(err)
			}

			for i := 0; i < countIntegers; i++ {
				num := rand.Int() * 100
				_, err = resTxt.WriteString(fmt.Sprintln(num))
				if err != nil {
					log.Fatal(err)
				}
			}
			err = resTxt.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}
	wg.Wait()
}
