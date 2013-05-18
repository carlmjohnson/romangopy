package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var NotRomanNumeral = errors.New("Input was not a valid roman numeral.")

var (
	re          = regexp.MustCompile("^(M*)(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$")
	conversions = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
)

func roman_to_arabic(input string) (int, error) {
	//The regexp matches by default
	if len(input) == 0 {
		return 0, NotRomanNumeral
	}

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
		//The first match is the whole thing. Ignore it.
		if i == 0 {
			continue
		}
		if len(match) == 2 {
			first, second := conversions[match[0]], conversions[match[1]]
			if second > first {
				total += second - first
			} else {
				total += second + first
			}
		} else {
			for j := range match {
				total += conversions[match[j]]
			}
		}
	}
	return total, nil
}

func arabic_to_roman(n int) string {
	return "I"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if n, err := strconv.Atoi(line); err == nil && n > 0 && n < 4000 {
			fmt.Println(arabic_to_roman(n))
		} else if roman, err := roman_to_arabic(line); err == nil {
			fmt.Println(roman)
		} else {
			fmt.Println(line)
		}
	}
}
