Go for Pythonistas

##Motivation
##Basics
| Python      |   Go        | Comment  |
|:-----------:|:-----------:|:---------|
`list`        |`[]T` ("slice")  |Both should really be called a "vector" or "dynamic array".
`tuple`       |`struct`      |Go also has multiple returns, which it uses instead of tuple returns.
`bool`        |`bool`        |Go uses "true" not "True" (which Python ought to also to be PEP-8 consistent but too late).
`None`        |`nil`         |Not a one-to-one replacement, but you know that already. Only slices, maps, channels, interfaces, and pointers can be `nil` in Go, but in practice that's enough.
`dict`        |`map[T1]T2`   |Both should really be called a "hash map".
`set`         |`map[T1]bool` |Go doesn't have a true set type, but a map to bool works as well, since unset keys default to false.
`str`/`unicode` |`string`      |Because the Go devs created UTF-8, Go is strongly biased towards UTF-8.
`bytes`/`str`   |`[]byte`      |Unlike Python `bytes` or Go `string`, `[]byte` is mutable.
-               |`rune`        |A `rune` is one Unicode codepoint. A `string` is both a list of `byte`s and a list of `rune`s.
generator    |`chan T`      | Channels and generators aren't the exact same, but in both cases, you're trying to make it easy to get information out of a coroutine in a synchronous way.
`Exception`  | `error`    | In Go, it's conventional for methods and functions to return a pair of `result, error` instead of raising exceptions. There is however a `panic` function that mostly works like an exception.
package      | package     | Python has both modules (files) and packages (folders).
module       | -          | Go automatically stitches all the files in a directory into one package.
`_name`      | `lowerCase`   | In Python, all module names except those starting with `_` are exported by default, but you can control this by setting `__all__`. In Go, uppercase names are exported and lowercase names are not.

##Example

##Conclusions

Original Python version:

	$ time python roman.py < test_input.txt > /dev/null
	
	real    0m5.874s
	user    0m5.847s
	sys 0m0.017s

Improved Python version:

	$ time python roman.py < test_input.txt > /dev/null
	
	real    0m2.023s
	user    0m2.007s
	sys 0m0.014s


Using standard regexp:

	$ time romangopy < test_input.txt > /dev/null
	
	real    0m1.486s
	user    0m1.340s
	sys 0m0.144s

Using rubex regexp:

	$ time romangopy < test_input.txt > /dev/null
	
	real    0m0.871s
	user    0m0.718s
	sys 0m0.146s
