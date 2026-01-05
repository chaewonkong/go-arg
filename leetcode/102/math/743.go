package math

import "strconv"

func fizzBuzz(n int) []string {
	const fizz = "Fizz"
	const buzz = "Buzz"
	const fizzBuzz = "FizzBuzz"

	ans := []string{}

	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			ans = append(ans, fizzBuzz)
			continue
		}

		if i%3 == 0 {
			ans = append(ans, fizz)
			continue
		}

		if i%5 == 0 {
			ans = append(ans, buzz)
			continue
		}

		ans = append(ans, strconv.Itoa(i))
	}

	return ans
}
