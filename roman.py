from re import search

conversions = {'I':1, 'V':5, 'X':10, 'L':50, 'C':100, 'D':500, 'M':1000}

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
    expanded = [int(tuple[1])*(10**(len(input)-tuple[0])) for tuple in enumerate(input, 1)]
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
    
while True:
    input = raw_input("Please enter an integer or roman numeral which you would like to be converted: ")
    if input.isdigit() and int(input)>=1: print "The roman numeral conversion of %s is: %s" %(input, arabic_to_roman(input))
    elif input.isalpha(): 
        try: print "The arabic integer conversion of %s is: %d" %(input, roman_to_arabic(input))
        except AttributeError: print "Sorry, your input was invalid."
    else: print "Sorry, your input was invalid."