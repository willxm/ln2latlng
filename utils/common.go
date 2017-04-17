package utils

import "strconv"

func Float2Str(fv float64) string {
	return strconv.FormatFloat(fv, 'f', 2, 64)
}
