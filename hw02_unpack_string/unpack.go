package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var alpha, digit rune
	var result string

	for _, r := range str {
		if string(r) == " " {
			return "", ErrInvalidString
		}
		if unicode.IsDigit(r) {
			digit = r
			if alpha == 0 {
				return "", ErrInvalidString
			}
			result += strings.Repeat(string(alpha), int(digit)-'0')
			alpha = 0
			continue
		} else if alpha != 0 {
			result += string(alpha)
		}
		alpha = r
	}
	if alpha != 0 {
		result += string(alpha)
	}
	return result, nil
}
