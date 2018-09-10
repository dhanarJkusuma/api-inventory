package helper

import (
	"strconv"
	"strings"
)

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 1, 64)
}

func IntToString(input_num int64) string {
	return strconv.FormatInt(input_num, 10)
}

func StringToInt64(s string) (int64, error) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func StringToFloat64(s string) (float64, error) {
	curr, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return curr, nil
}

func GetOrderIDByOutcomingProductNote(note string) string {
	raw := strings.Split(note, " ")
	if len(raw) > 1 {
		return raw[1]
	}
	return ""
}
