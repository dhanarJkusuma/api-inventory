package helper

import "strconv"

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 1, 64)
}

func IntToString(input_num int64) string {
	return strconv.FormatInt(input_num, 10)
}
