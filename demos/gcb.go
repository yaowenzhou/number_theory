package demos

import "golang.org/x/exp/constraints"

// Gcb 辗转相除法求最大公约数
func Gcb[T constraints.Integer](i, j T) T {
	for i != 0 {
		i, j = j%i, i
	}
	return j
}
