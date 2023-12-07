package common

func SumArray(arr []int) int {
	total := 0
	for _, num := range arr {
		total += num
	}
	return total
}
