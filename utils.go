package utils

func Contains(stringSlice []string, stringToFind string) bool {
	for _, str := range stringSlice {
		if str == stringToFind {
			return true
		}
	}
	return false
}

func ContainsInt(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
