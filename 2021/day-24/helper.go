package main

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func helper(input [14]int) int {
	lastConsumedInputIdx := 0
	w, x, y, z := 0, 0, 0, 0

	// make sure compiler doesn't complain
	w = w * x * y
	lastConsumedInputIdx = lastConsumedInputIdx * lastConsumedInputIdx

	return z
}
