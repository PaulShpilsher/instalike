package utils

import (
	"strconv"
	"strings"
)

// ParseStringToInt64Array function - postgres returns an integer ARRAY column as a string in format "{1,2,3,4}".
// this function parses this string and returns []int64 data
func ParseStringToInt64Array(array string) []int64 {

	if array == "" || array == "{}" || array[0] != '{' || array[len(array)-1] != '}' {
		return []int64{}
	}

	toParse := array[1 : len(array)-1]
	parts := strings.Split(toParse, ",")
	return Map(parts, func(s string) int64 {
		v, _ := strconv.ParseInt(s, 10, 64)
		return v
	})

}
