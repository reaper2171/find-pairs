package utils

func TwoSum(numbers []int, target int) [][]int {
	// taking mutli-map here with []int as key's value here coz duplicate indices may satify the target value
	indMap := make(map[int][]int)
	var solutions [][]int

	for i, num := range numbers {
		rem := target - num
		if indices, ok := indMap[rem]; ok {
			for _, index := range indices {
				solutions = append(solutions, []int{index, i})
			}
		}
		indMap[num] = append(indMap[num], i)
	}
	return solutions
}
