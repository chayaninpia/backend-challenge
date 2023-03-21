package src

import (
	"strconv"
	"strings"
)

func LeftRightEqual(input string) string {
	countL := strings.Count(input, "L") + 1
	finalResult := ""
	for !strings.Contains(finalResult, "0") {
		countL--
		startDigit := countL
		dst := make([]byte, 0)
		for k, v := range input {
			switch v {
			case 'L':
				if k == 0 {
					startDigit++
				}
				startDigit--
				dst = strconv.AppendInt(dst, int64(startDigit), 10)
			case 'R':
				if k == 0 {
					startDigit--
				}
				startDigit++
				dst = strconv.AppendInt(dst, int64(startDigit), 10)
			case '=':
				dst = strconv.AppendInt(dst, int64(startDigit), 10)
			}
		}
		finalResult = string(dst)
	}
	return finalResult
}
