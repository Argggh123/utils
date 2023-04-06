package utils

func InSlice(stringSlice []string, stringToFind string) bool {
	for _, str := range stringSlice {
		if str == stringToFind {
			return true
		}
	}
	return false
}

func InSliceInt(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
