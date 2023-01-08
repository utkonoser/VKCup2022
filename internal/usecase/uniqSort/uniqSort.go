package uniqSort

import (
	"goElimination/internal/usecase/sort"
	"goElimination/internal/usecase/uniq"
)

// SetVar - множество уникальных значений
var SetVar uniq.SetStruct

// RunUniqSort - основная функция кейса с сортировкой уникальных значений
func RunUniqSort(path string) {
	SetVar.RunReadAllUniq(path)
	items := SetVar.Items().ShowItems()
	sort.QSort(items)
	sort.CreateTxt(items, path)
}
