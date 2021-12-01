package goalgos

//SumNumbersInList to sum all numbers in list
func SumNumbersInList(list []int) int {
	var sum int
	for _, i := range list {
		sum = sum + i
	}
	return sum
}
