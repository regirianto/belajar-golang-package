package utils

func Sum(scores []int) (total int) {
	// var total int
	for _, score := range scores {
		total += score
	}
	return
}
