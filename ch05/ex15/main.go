package main

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

func max(vals ...int) int {
	m := minInt
	for _, v := range vals {
		if v > m {
			m = v
		}
	}
	return m
}

func min(vals ...int) int {
	m := maxInt
	for _, v := range vals {
		if v < m {
			m = v
		}
	}
	return m
}

func safeMax(val int, vals ...int) int {
	m := val
	for _, v := range vals {
		if v > m {
			m = v
		}
	}
	return m
}

func safeMin(val int, vals ...int) int {
	m := val
	for _, v := range vals {
		if v < m {
			m = v
		}
	}
	return m
}
