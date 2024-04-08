package sum

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numberToSum ...[]int) []int {
	// lengthOfNumbers := len(numberToSum)
	// sums = make([]int, lengthOfNumbers)

	// for i, numbers := range numberToSum {
	// 	sums[i] = Sum(numbers)
	// }

	var sums []int
	for _, number := range numberToSum {
		sums = append(sums, Sum(number))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
				tail := numbers[1:]
				sums = append(sums, Sum(tail))
		}
	}
	return sums
}
