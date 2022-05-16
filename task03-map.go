package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	i := []int{}
	for k, _ := range input {
		i = append(i, k)
	}
	sort.Ints(i)
	for _, v := range i {
		result = append(result, input[v])
	}
	return result
}
