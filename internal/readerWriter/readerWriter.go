package rw

import (
	"bufio"
	"fmt"
	"os"
)

func Read() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter example: ")
	text, _ := reader.ReadString('\n')
	return text
}

func Write(s string) {
	fmt.Print()
}
