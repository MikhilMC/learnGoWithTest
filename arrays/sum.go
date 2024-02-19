package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumOrdinaryForLoop(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sum := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sum[i] = Sum(numbers)
	}

	return sum
}
