package util

import "strconv"

func Str2Uint64(s string) uint64 {
	if i, err := strconv.ParseUint(s, 10, 64); err == nil {
		return i
	}
	return 0
}

func Str2int(s string) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	return 0
}
