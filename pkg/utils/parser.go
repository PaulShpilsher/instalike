package utils

import (
	"strconv"
	"strings"
)

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
