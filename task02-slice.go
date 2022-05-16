package homework

func reverse(input []int64) (result []int64) {
	for i := 1; i < len(input); i++ {
		result = append(result, input[len(input)-i])
	}
	return result
}
