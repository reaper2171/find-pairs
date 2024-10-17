package main

func TwoSum(numbers []int, target int) [][]int {
	indMap := make(map[int]int)
	var solutions [][]int

	for i, num := range numbers {
		rem := target - num
		if index, ok := indMap[rem]; ok {
			solutions = append(solutions, []int{index, i})
		}
		indMap[num] = i
	}
	return solutions
}
