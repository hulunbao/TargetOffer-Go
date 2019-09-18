package problem11

import "errors"

// Pow 函数
func Pow(base float64, exp int) (float64, error) {
	if exp == 0 {
		return 1, nil
	}
	if base == 0 && exp < 0 {
		return -1, errors.New("base == 0 and exp < 0 ")
	}
	if exp > 0 {
		return PowNormal(base, exp), nil
	} else {
		res := PowNormal(base, -exp)
		res = 1.0 / res
		return res, nil
	}
}

// PowNormal 取exp次方
func PowNormal(base float64, exp int) float64 {
	res, temp := 1.0, base
	for exp != 0 {
		if exp&1 == 1 {
			res *= temp
		}
		temp *= temp
		exp >>= 1
	}
	return res
}
