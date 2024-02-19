package iteration

import (
	"fmt"
	"strings"
)

func Repeat(character string, repeatedCount int) string {
	return strings.Repeat(character, repeatedCount)
}

func RepeatIterable(character string, repeatedCount int) string {
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
