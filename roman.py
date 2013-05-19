"""A simple module to convert arabic numerals to roman numerals and vice
versa. Taken from http://pastebin.com/V6jHJZX3."""

from re import search

conversions = {'I':1, 'V':5, 'X':10, 'L':50, 'C':100, 'D':500, 'M':1000}
reverse_conversions = {
        0: "", 1: "I", 2: "II", 3: "III", 4: "IV",
        5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX",
        10: "X", 20: "XX", 30: "XXX", 40: "XL",
        50: "L", 60: "LX", 70: "LXX", 80: "LXXX", 90: "XC",
        100: "C", 200: "CC", 300: "CCC", 400: "CD",
        500: "D", 600: "DC", 700: "DCC", 800: "DCCC", 900: "CM",
}

def roman_to_arabic(input):
    groups = search("^(M*)(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$", input.upper()).groups()
    total = 0
    for group in groups:
        if len(group)==2 and conversions[group[1]]>conversions[group[0]]:
            total += (conversions[group[1]]-conversions[group[0]])
        elif group:
            for digit in group:
                total += conversions[digit]
    return total

def arabic_to_roman(input):
    roman = []
    
    n = int(input)
    if n < 1:
        raise ValueError("Only integers > 1 can be put in roman numerals")
    #Get thousands digit
    #Put that many Ms upfront.
    thousands = n // 1000
    roman.append("M" * thousands)

    #Chop off remaining thousands.
    n %= 1000

    #Get hundreds digit, use dict to look that up.
    
    hundreds = (n // 100) * 100
    roman.append(reverse_conversions[hundreds])

    #Chop off remaining hundreds.
    n %= 100

    #Get tens digit
    tens = (n // 10) * 10
    roman.append(reverse_conversions[tens])

    #Get ones digit
    ones = n % 10
    roman.append(reverse_conversions[ones])

    return "".join(roman)


def original_arabic_to_roman(input):
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

def interactive_main():
    while True:
        input = raw_input("Please enter an integer or roman numeral which you would like to be converted: ")
        if input.isdigit() and int(input)>=1: print "The roman numeral conversion of %s is: %s" %(input, arabic_to_roman(input))
        elif input.isalpha(): 
            try: print "The arabic integer conversion of %s is: %d" %(input, roman_to_arabic(input))
            except AttributeError: print "Sorry, your input was invalid."
        else: print "Sorry, your input was invalid."

def generate_tests():
    import random
    
    funcs = [arabic_to_roman, lambda s: s]
    
    while True:
        print random.choice(funcs)(str(random.randint(1,3999)))

def std_main():
    import fileinput
    
    for input in fileinput.input():
        input = input.strip()
        if input.isdigit() and int(input)>=1: 
            print arabic_to_roman(input)
        elif input.isalpha(): 
            try: 
                print roman_to_arabic(input)
            except AttributeError: 
                print input
        else:
            print input

if __name__ == "__main__":
    std_main()
