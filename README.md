# calculator
The calculator can perform addition, subtraction, multiplication, and division with two numbers: a + b, a - b, a * b, a / b. Data is entered in a single line (see example below). Solutions where each number and the arithmetic operation are entered on a new line are considered incorrect.

The calculator can work with both Arabic (1, 2, 3, 4, 5...) and Roman (I, II, III, IV, V...) numbers.

The calculator should accept input numbers from 1 to 10 inclusive. The output numbers are not restricted in size and can be any value.

The calculator can only work with integers.

The calculator can only operate with Arabic or Roman numerals at the same time. If the user enters a line like 3 + II, the calculator should trigger an error and halt operation.

When using Roman numbers, the output should be in Roman numerals, and for Arabic numbers, the output is expected to be in Arabic.

If the user inputs unsuitable numbers, the application triggers an error and terminates.

If the user enters a line that does not correspond to one of the aforementioned arithmetic operations, the application triggers an error and terminates.

The result of the division operation is an integer, with the remainder discarded.

The result of calculations with Arabic numbers can be negative or zero. For calculations with Roman numerals, only positive results are allowed; if the result is less than one, the program should trigger an error.