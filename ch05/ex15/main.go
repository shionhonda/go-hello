package main

import "fmt"

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

func safeMax(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("no values provided")
	}
	m := minInt
	for _, v := range vals {
		if v > m {
			m = v
		}
	}
	return m, nil
}

func safeMin(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("no values provided")
	}
	m := maxInt
	for _, v := range vals {
		if v < m {
			m = v
		}
	}
	return m, nil
}
