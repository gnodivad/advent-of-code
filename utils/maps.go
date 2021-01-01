package utils

import (
	"sort"
	"strings"
)

type Pair struct {
	Key   string
	Value string
}

type PairList []Pair

func SortByValue(maps map[string]string) PairList {
	kvList := make([]Pair, 0)

	for k, v := range maps {
		kvList = append(kvList, Pair{k, v})
	}

	sort.Slice(kvList, func(i, j int) bool {
		return strings.Compare(kvList[i].Value, kvList[j].Value) == -1
	})

	return kvList
}
