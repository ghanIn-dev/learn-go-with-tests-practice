package main

func main() {

}

// func Sum(numbers [5]int) int {
// 	sum := 0
// 	// for i := 0; i < 5; i++ {
// 	// 	sum += numbers[i]
// 	// }

// 	//refactor
// 	for _, number := range numbers {
// 		sum += number
// 	}
// 	return sum
// }

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)

	// for i, numbers := range numbersToSum {
	// 	sums[i] = Sum(numbers)
	// }

	//위 코드가 아래와 같이 리팩토링 되면서 slice의 용량에 대한 걱정이 없어진다.
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tails := numbers[1:]
			sums = append(sums, Sum(tails))
		}
	}
	return sums
}
