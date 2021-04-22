package utils

func ReverseInt(n int) (res int) {
	for n != 0 {
		res = res*10 + n%10
		if res > 2147483647 || res < -2147483648 {
			return 0
		}

		n /= 10
	}

	return
}
