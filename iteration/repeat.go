package iteration

import "fmt"

func Repeat(character string, repeatedCount int) string {
	var repeated string

	for i := 0; i < repeatedCount; i++ {
		repeated += character
	}

	return repeated
}

func ExampleRepeat() {
	character := "a"
	count := 6
	fmt.Println(Repeat(character, count))
}
