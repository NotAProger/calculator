package main

import (
	calc "calculator/internal/calculator"
	rw "calculator/internal/readerWriter"
)

func main() {
	rw.Write(calc.Run(rw.Read()))
}
