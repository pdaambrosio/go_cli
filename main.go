package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// The main function reads user input, trims whitespace, converts it to uppercase, and then prints it.
func main () {
	fmt.Println("What would you like me to scean?")
	in := bufio.NewReader(os.Stdin)
	inputRead, _ := in.ReadString('\n')
	inputRead = strings.TrimSpace(inputRead)
	inputRead = strings.ToUpper(inputRead)

	fmt.Println(inputRead)
}
