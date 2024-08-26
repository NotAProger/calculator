# calculator

# eng
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

Program example:
```
Input: 
1 + 2 
Output: 
3

Input: 
VI / III 
Output: 
II

Input:
I - II 
Output: 
Panic, as there are no negative numbers in the Roman numeral system.

Input: 
I + 1 
Output: 
Panic, as different numeral systems are being used simultaneously.

Input: 
1 
Output: 
Panic, as the line is not a mathematical operation.

Input: 
1 + 2 + 3 
Output: 
Panic, as the operation format does not meet the requirements—two operands and one operator (+, -, /, *).
```

# ru

Калькулятор умеет выполнять операции сложения, вычитания, умножения и деления с двумя числами: a + b, a - b, a * b, a / b. Данные передаются в одну строку (смотри пример ниже). Решения, в которых каждое число и арифметическая операция передаются с новой строки, считаются неверными.

Калькулятор умеет работать как с арабскими (1, 2, 3, 4, 5…), так и с римскими (I, II, III, IV, V…) числами.

Калькулятор должен принимать на вход числа от 1 до 10 включительно, не более. На выходе числа не ограничиваются по величине и могут быть любыми.

Калькулятор умеет работать только с целыми числами.

Калькулятор умеет работать только с арабскими или римскими цифрами одновременно, при вводе пользователем строки вроде 3 + II калькулятор должен выдать панику и прекратить работу.

При вводе римских чисел ответ должен быть выведен римскими цифрами, соответственно, при вводе арабских — ответ ожидается арабскими.

При вводе пользователем не подходящих чисел приложение выдаёт панику и завершает работу.

При вводе пользователем строки, не соответствующей одной из вышеописанных арифметических операций, приложение выдаёт панику и завершает работу.

Результатом операции деления является целое число, остаток отбрасывается.

Результатом работы калькулятора с арабскими числами могут быть отрицательные числа и ноль. Результатом работы калькулятора с римскими числами могут быть только положительные числа, если результат работы меньше единицы, программа должна выдать панику.

Пример работы программы:
```
Input:
1 + 2
Output:
3

Input:
VI / III
Output:
II

Input:
I - II
Output:
Выдача паники, так как в римской системе нет отрицательных чисел.

Input:
I + 1
Output:
Выдача паники, так как используются одновременно разные системы счисления.

Input:
1
Output:
Выдача паники, так как строка не является математической операцией.

Input:
1 + 2 + 3
Output:
Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).
```