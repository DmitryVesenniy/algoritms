package main

func MaxElement(vals []int, left int, right int) int {
	if (right - left) == 1 {
		return vals[left]
	}

	//делим массив на две части
	var mid int = (left + right) / 2

	//Рекурсия
	var max1 int = MaxElement(vals, left, mid)
	var max2 int = MaxElement(vals, mid, right)

	if max1 > max2 {
		return max1
	}
	return max2
}
