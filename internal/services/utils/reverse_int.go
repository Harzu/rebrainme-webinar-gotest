package utils

func reverseInt(num int) int {
	var result, prevResult int
	for num != 0 {
		digit := num % 10
		result = result*10 + digit

		if (result-digit)/10 != prevResult {
			return 0
		}

		prevResult = result
		num /= 10
	}
	return result
}
