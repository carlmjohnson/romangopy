package main

import (
	"fmt"
	"github.com/earthboundkid/stdin"
	_ "regexp"
	"strconv"
)

func roman_to_arabic(input string) (int, error) {
	return 1, nil
}

func arabic_to_roman(n int) string {
	return "I"
}

func main() {
	for line := range stdin.Chan() {
		if n, err := strconv.Atoi(line); err == nil && n > 0 && n < 4000 {
			fmt.Println(arabic_to_roman(n))
		} else if roman, err := roman_to_arabic(line); err == nil {
			fmt.Println(roman)
		} else {
			fmt.Println(line)
		}
	}
}
