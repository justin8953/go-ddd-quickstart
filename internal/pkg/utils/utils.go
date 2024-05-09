package utils

import "strconv"

func ConvertStringToInt(s string) (int, error) {
	if s == "" {
		return 0, nil
	}
	return strconv.Atoi(s)
}

func ConvertBytesToString(b []byte) string {
	return string(b)
}
