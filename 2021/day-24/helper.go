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

	w = input[lastConsumedInputIdx]
	lastConsumedInputIdx++

	x = z%26 + 13
	x = btoi(x != w)
	y = 25*x + 1
	z *= y
	y = w + 13
	y *= x
	z += y

	w = input[lastConsumedInputIdx]
	lastConsumedInputIdx++

	x = z%26 + 11
	x = btoi(x != w)
	y = 25*x + 1
	z *= y
	y = w + 10
	y *= x
	z += y

	w = input[lastConsumedInputIdx]
	lastConsumedInputIdx++

	x = z%26 + 15
	x = btoi(x != w)
	y = 25*x + 1
	z *= y
	y = w + 5
	y *= x
	z += y

	w = input[lastConsumedInputIdx]
	lastConsumedInputIdx++

	x = z%26 - 11
	x = btoi(x != w)
	y = 25*x + 1
	z /= 26
	z *= y
	y = w + 14
	y *= x
	z += y

	w = input[lastConsumedInputIdx]
	lastConsumedInputIdx++

	x = z%26 + 14
	x = btoi(x != w)
	y = 25*x + 1
	z *= y
	y = w + 5
	y *= x
	z += y

	w = input[lastConsumedInputIdx]
	lastConsumedInputIdx++

	x = z % 26
	x = btoi(x != w)
	y *= 0
	y = 25*x + 1
	z /= 26
	z *= y
	y = w + 15
	y *= x
	z += y

	w = input[lastConsumedInputIdx]
	lastConsumedInputIdx++

	x = z%26 + 12
	x = btoi(x != w)
	y *= 0
	y = 25*x + 1
	z *= y
	y = w + 4
	y *= x
	z += y

	w = input[lastConsumedInputIdx]
	lastConsumedInputIdx++

	x = z%26 + 12
	x = btoi(x != w)
	y = 25*x + 1
	z *= y
	y = w + 11
	y *= x
	z += y

	w = input[lastConsumedInputIdx]
	lastConsumedInputIdx++

	x = z%26 + 14
	x = btoi(x != w)
	y = 25*x + 1
	z *= y
	y = w + 1
	y *= x
	z += y

	w = input[lastConsumedInputIdx]
	lastConsumedInputIdx++

	x = z%26 - 6
	x = btoi(x != w)
	y = 25*x + 1
	z /= 26
	z *= y
	y = w + 15
	y *= x
	z += y

	w = input[lastConsumedInputIdx]
	lastConsumedInputIdx++

	x = z%26 - 10
	x = btoi(x != w)
	y = 25*x + 1
	z /= 26
	z *= y
	y = w + 12
	y *= x
	z += y

	w = input[lastConsumedInputIdx]
	lastConsumedInputIdx++

	x = z%26 - 12
	x = btoi(x != w)
	y = 25*x + 1
	z /= 26
	z *= y
	y = w + 8
	y *= x
	z += y

	w = input[lastConsumedInputIdx]
	lastConsumedInputIdx++

	x = z%26 - 3
	x = btoi(x != w)
	y = 25*x + 1
	z /= 26
	z *= y
	y = w + 14
	y *= x
	z += y

	w = input[lastConsumedInputIdx]
	lastConsumedInputIdx++

	x = z%26 - 5
	x = btoi(x != w)
	y = 25*x + 1
	z /= 26
	z *= y
	y = w + 9
	y *= x
	z += y

	return z
}
