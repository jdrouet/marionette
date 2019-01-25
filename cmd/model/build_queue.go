package model

func Omit(queue []string, value string) []string {
	result := []string{}
	for _, item := range queue {
		if value != item {
			result = append(result, item)
		}
	}
	return result
}

func Contains(queue []string, value string) bool {
	for _, item := range queue {
		if value == item {
			return true
		}
	}
	return false
}

func RemoveDuplicates(queue []string) []string {
	result := []string{}
	for _, item := range queue {
		if !Contains(result, item) {
			result = append(result, item)
		}
	}
	return result
}
