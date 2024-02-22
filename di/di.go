package main

import (
	"bytes"
	"fmt"
	"os"
)

func Greet(writer *bytes.Buffer, name string) {
	fmt.Printf("Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Elodie")
}
