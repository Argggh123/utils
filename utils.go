package utils

func Contains(stringSlice []string, stringToFind string) bool {
	for _, str := range stringSlice {
		if str == stringToFind {
			return true
		}
	}

	return false
}
