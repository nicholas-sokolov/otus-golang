package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var alpha, digit rune
	var result strings.Builder

	for _, r := range str {
		if string(r) == " " {
			return "", ErrInvalidString
		}
		if unicode.IsDigit(r) {
			digit = r
			if alpha == 0 {
				return "", ErrInvalidString
			}
			for i := 0; i < int(digit)-'0'; i++ {
				result.WriteRune(alpha)
			}
			alpha = 0
			continue
		} else if alpha != 0 {
			result.WriteRune(alpha)
		}
		alpha = r
	}
	if alpha != 0 {
		result.WriteRune(alpha)
	}
	return result.String(), nil
}
