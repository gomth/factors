package factors

import (
	"errors"

	"github.com/gomth/primes"
)

// GetUsingTrialDivisionForInt - if val is prime number already, return value will be []int{val}, error, indicating an error
func GetUsingTrialDivisionForInt(val int) ([]int, error) {
	primeNumers, err := primes.Generate(int(val))
	if err != nil {
		return nil, err
	}

	// if prime number was passed
	if primeNumers[len(primeNumers)-1] == val {
		return primeNumers[len(primeNumers)-1:], errors.New("Prime number found")
	}

	var result []int

	valLeft := val
	for _, prime := range primeNumers {

		for {
			if valLeft%prime == 0 {
				result = append(result, prime)
				valLeft = valLeft / prime
			} else {
				break
			}

			if valLeft <= 1 {
				break
			}
		}
	}

	return result, nil
}
