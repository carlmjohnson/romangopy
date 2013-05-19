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
	re = regexp.MustCompile("^(M*)(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$")

	conversions = map[byte]int{
		'I': 1, 'V': 5, 'X': 10,
		'L': 50, 'C': 100, 'D': 500,
		'M': 1000,
	}
	rConversions = map[int]string{
		1: "I", 2: "II", 3: "III", 4: "IV",
		5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX",
		10: "X", 20: "XX", 30: "XXX", 40: "XL",
		50: "L", 60: "LX", 70: "LXX", 80: "LXXX", 90: "XC",
		100: "C", 200: "CC", 300: "CCC", 400: "CD",
		500: "D", 600: "DC", 700: "DCC", 800: "DCCC", 900: "CM",
	}

	romanBacking = [4]string{}
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

/*
   expanded = [int(s)*(10**(len(input)-n)) for n, s in enumerate(input, 1)]
   roman = []
   for count, number in enumerate(expanded):
       if number not in conversions.values():
           if number>1000:
               expanded[count] = 1000
               for time in range((number/1000)-1):
                   expanded.insert(1, 1000)
           elif 1<int(str(number)[0])<4:
               expanded[count] = 10**(len(str(number))-1)
               for time in range(int(str(number)[0])-1):
                   expanded.insert(count+1, expanded[count])
           elif int(str(number)[0])==4:
               expanded[count] = 10**(len(str(number))-1)
               expanded.insert(count+1, 5*(expanded[count]))
           elif 5<int(str(number)[0])<9:
               expanded[count] = 5*(10**(len(str(number))-1))
               for time in range(int(str(number)[0])-5):
                   expanded.insert(count+1, 10**(len(str(number))-1))
           elif int(str(number)[0])==9:
               expanded[count] = 10**(len(str(number))-1)
               expanded.insert(count+1, 10*(expanded[count]))
   for number in expanded:
       for entry in conversions:
           if conversions[entry]==number:
               roman.append(entry)
   return "".join(roman)
*/
func arabic_to_roman(n int) string {
	roman := romanBacking[:0]

	//Get thousands digit
	//Put that many Ms upfront.
	if thousands := n / 1000; thousands != 0 {
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if n, err := strconv.Atoi(line); err == nil && n > 0 {
			fmt.Println(arabic_to_roman(n))
		} else if roman, err := roman_to_arabic(line); err == nil {
			fmt.Println(roman)
		} else {
			fmt.Println(line)
		}
	}
}
