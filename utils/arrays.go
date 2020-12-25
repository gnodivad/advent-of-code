package utils

func Contains(arr []int, target int) bool {
	for _, a := range arr {
		if a == target {
			return true
		}
	}
	return false
}
