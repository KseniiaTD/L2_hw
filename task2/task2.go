package task2

import (
	"fmt"
	"unicode"
)

func addElem(repeatCnt int, elemToAdd rune, res []rune) []rune {
	if repeatCnt >= 0 {
		for l := 0; l < repeatCnt; l++ {
			res = append(res, elemToAdd)
		}
	} else {
		res = append(res, elemToAdd)
	}
	return res
}

func Unpack(str string) (string, error) {
	res := []rune{}
	repeatCnt := -1
	isEscape := false
	escape := rune('\u005c')
	var elemToAdd rune

	if len(str) == 0 {
		return "", nil
	}

	for i, elem := range str {

		if unicode.IsDigit(elem) && i == 0 {
			return "", fmt.Errorf("invalid string: %s", str)
		}

		if isEscape {
			if elemToAdd != 0 {
				res = addElem(repeatCnt, elemToAdd, res)
			}
			elemToAdd = elem
			isEscape = false
			continue
		}

		if elem == escape {
			isEscape = true
			continue
		}

		if unicode.IsDigit(elem) {
			if repeatCnt == -1 {
				repeatCnt = int(elem - '0')
			} else {
				repeatCnt = repeatCnt*10 + int(elem-'0')
			}
			continue
		}

		if elemToAdd != 0 {
			res = addElem(repeatCnt, elemToAdd, res)
		}
		repeatCnt = -1
		elemToAdd = elem

	}

	if isEscape {
		return "", fmt.Errorf("invalid string: %s", str)
	}

	res = addElem(repeatCnt, elemToAdd, res)

	return string(res), nil
}
