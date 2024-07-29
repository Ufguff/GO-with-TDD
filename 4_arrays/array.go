package arrays

func Sum(numbers []int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	var result []int

	for _, v := range numbers {
		result = append(result, Sum(v))
	}
	return result
}

func SumAllTails(numbers ...[]int) []int {
	var result []int

	for _, v := range numbers {
		if len(v) == 0 {
			result = append(result, 0)
		} else {
			result = append(result, Sum(v[1:]))
		}

	}
	return result
}
