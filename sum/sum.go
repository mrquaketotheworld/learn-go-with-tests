package sum

func Sum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	sumSlice := make([]int, 0, len(numbers))
	for _, innerNumbers := range numbers {
		sumSlice = append(sumSlice, Sum(innerNumbers))
	}
	return sumSlice
}

func SumAllTails(numbers ...[]int) []int {
	sumSlice := make([]int, 0, len(numbers))
	for _, innerNumbers := range numbers {
		if len(innerNumbers) == 0 {
			sumSlice = append(sumSlice, 0)
		} else {
			sumSlice = append(sumSlice, Sum(innerNumbers[1:]))
		}
	}
	return sumSlice
}
