package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var result strings.Builder
	k := false

	for i, j := 0, 1; i < len(s); i, j = i+1, j+1 {
		if !k {
			_, err := strconv.Atoi(string(s[i]))
			if err == nil {
				return "", ErrInvalidString
			}
		}

		b := string(s[i])

		if b == "\\" && !k {
			k = true
			continue
		}

		if j < len(s) {
			num, err := strconv.Atoi(string(s[j]))

			if err == nil {
				repstr := strings.Repeat(b, num)
				result.WriteString(repstr)
				i++
				j++
			} else {
				result.WriteString(string(s[i]))
			}
		} else {
			result.WriteString(string(s[i]))
		}

		k = false
	}

	return result.String(), nil
}
