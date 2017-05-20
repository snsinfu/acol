package utils

/*
IntMin returns the minimum of given two integers.
*/
func IntMin(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

/*
IntMax returns the maximum of given two integers.
*/
func IntMax(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

/*
IntMaxReduce returns the maximum integer in given sequence and given number.
*/
func IntMaxReduce(xs []int, init int) int {
	max := init
	for _, x := range xs {
		max = IntMax(max, x)
	}
	return max
}

/*
IntSum computes the sum of given integers. Returns zero if input is empty.
*/
func IntSum(xs []int) int {
	sum := 0
	for _, x := range xs {
		sum += x
	}
	return sum
}
