package utils

func Pow(n, p int) int {
	if p == 0 {
		return 1
	}
	base := n
	for i := 0; i < p-1; i++ {
		n *= base
	}
	return n
}
