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

//This is the error returned when func RomanToArabic fails.
var NotRomanNumeral = errors.New("Input was not a valid roman numeral.")

var (
	re = regexp.MustCompile("^(M*)(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$")

	conversions = map[byte]int{
		'I': 1, 'V': 5, 'X': 10,
		'L': 50, 'C': 100, 'D': 500,
		'M': 1000,
	}
	rConversions = map[uint]string{
		1: "I", 2: "II", 3: "III", 4: "IV",
		5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX",
		10: "X", 20: "XX", 30: "XXX", 40: "XL",
		50: "L", 60: "LX", 70: "LXX", 80: "LXXX", 90: "XC",
		100: "C", 200: "CC", 300: "CCC", 400: "CD",
		500: "D", 600: "DC", 700: "DCC", 800: "DCCC", 900: "CM",
	}
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if n, err := strconv.Atoi(line); err == nil && n > 0 {
			fmt.Println(ArabicToRoman(uint(n)))
		} else if roman, err := RomanToArabic(line); err == nil {
			fmt.Println(roman)
		} else {
			fmt.Println(line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

//RomanToArabic takes an input string and returns an int of that string
//as a roman numeral if possible. If input is not a valid roman numeral,
//it returns the error NotRomanNumeral.
func RomanToArabic(input string) (int, error) {
	//The regexp matches by default, so cut it off here.
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

//AbrabicToRoman takes an unsigned int (no negative numbers!) and 
//returns a string with that int as a roman numeral. If the number is
//more than three thousand, it just jams a lot of Ms on the front. Given
//zero it returns the empty string.
func ArabicToRoman(n uint) string {
	roman := make([]string, 0, 4)

	//Get thousands digit
	//Put that many Ms upfront.
	if thousands := int(n) / 1000; thousands != 0 {
		roman = append(roman, strings.Repeat("M", thousands))
	}

	//Chop off remaining thousands.
	n %= 1000

	//Get hundreds digit, use map to look that up.
	if hundreds := n / 100 * 100; hundreds != 0 {
		roman = append(roman, rConversions[hundreds])
	}

	//Chop off remaining hundreds.
	n %= 100

	//Get tens digit, use map to look that up.
	if tens := n / 10 * 10; tens != 0 {
		roman = append(roman, rConversions[tens])
	}

	//Get ones digit, etc.
	if ones := n % 10; ones != 0 {
		roman = append(roman, rConversions[ones])
	}

	return strings.Join(roman, "")
}
