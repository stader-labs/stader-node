package arr_utils

func ElementExistsInNumArray(arr []int64, element int64) bool {
	for _, v := range arr {
		if v == element {
			return true
		}
	}
	return false
}
