package dnquestions

// RepeatedString returns the number of times the letter 'a' appears in the string
// assuming that `str` repeated `n` times.
func RepeatedString(str string, n int) int {
	strLen := len(str)

	if strLen == 0 {
		return 0
	}

	count := 0
	for _, char := range str {
		if char == 'a' {
			count++
		}
	}

	count *= n / strLen
	for _, val := range str[:n%strLen] {
		if val == 'a' {
			count++
		}
	}

	return count
}

// MinimumRemovals returns the minimum number of removals needed to make the string
// alternating between 'A' and 'B'.
func MinimumRemovals(str string) int {
	count := 0

	for i := range str {
		if i == len(str)-1 {
			break
		}

		if str[i] == str[i+1] {
			count++
			i++
		}
	}

	return count
}

// MinSwapsToSort returns the minimum number of swaps needed to sort the array.
func MinSwapsToSort(arr []int) int {
	swaps := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] == i+1 {
			continue
		}

		temp := arr[i]
		arr[i] = arr[temp-1]
		arr[temp-1] = temp
		swaps++
		i--
	}

	return swaps
}
