package main

import (
	"errors"
	"fmt"
	"github.com/earthboundkid/stdin"
	"regexp"
	"strconv"
	"strings"
)

/*
   groups = search("^(M*)(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$", input.upper()).groups()
   total = 0
   for group in groups:
       if len(group)==2 and conversions[group[1]]>conversions[group[0]]:
           total += (conversions[group[1]]-conversions[group[0]])
       elif group:
           for digit in group:
               total += conversions[digit]
   return total
*/
var NotRomanNumeral = errors.New("Input was not a valid roman numeral.")

var re = regexp.MustCompile("^(M*)(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$")

func convert(digit byte) int {
	switch digit {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	panic("Only use func convert with valid roman numeral digits.")
}

func roman_to_arabic(input string) (int, error) {
	//Normalize input
	input = strings.ToUpper(input)

	//Find all matches
	matchlist := re.FindAllStringSubmatch(input, -1)

	//If it doesn't have a match, it's not a roman numeral
	if len(matchlist) == 0 {
		return 0, NotRomanNumeral
	}

	total := 0

	for i, match := range matchlist[0] {
		if i == 0 {
			continue
		}
		if len(match) == 2 && (convert(match[1]) > convert(match[0])) {
			total += convert(match[1]) - convert(match[0])
		} else if len(match) > 0 {
			for j := range match {
				total += convert(match[j])
			}
		}
	}
	return total, nil
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
