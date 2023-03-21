package src

import (
	"encoding/json"
	"os"
	"sort"
)

func MaxPath(path string) (int, error) {

	maxPath, err := readFile(path)
	if err != nil {
		return 0, err
	}
	var total int
	for _, v := range maxPath {
		if len(v) != 0 {
			sort.Slice(v, func(i, j int) bool {
				return v[i] > v[j]
			})
			total += v[0]
		}
	}

	return total, nil
}

func readFile(path string) ([][]int, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	maxPath := make([][]int, 0)
	if err := json.Unmarshal(b, &maxPath); err != nil {
		return nil, err
	}
	return maxPath, nil
}
