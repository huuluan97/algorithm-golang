package rsa

import (
	"math"
	"math/rand"
)

func prime(limit int) (primes []int) {
	sqrtLimit := int(math.Ceil(math.Sqrt(float64(limit))))
	exit := false
	primes = append(primes, 2, 3, 5)
	lastIndex := 2
	for primes[lastIndex] < sqrtLimit {
		if exit {
			break
		}
		for i := primes[lastIndex] + 2; i < primes[lastIndex]*primes[lastIndex]; i += 2 {
			found := true
			for _, v := range primes {
				if i%v == 0 {
					found = false
					break
				}
			}
			if found {
				primes = append(primes, i)
				lastIndex++
				if i >= sqrtLimit {
					exit = true
					break
				}
			}
		}
	}
	return
}

func generatePrimes(limit int) int {

	/*
		generate primes by factoring
		relies on the 30k+i, though better formulae exist
		where k >=0 and i = (1,7,11,13,17,13,19,23,29)
	*/
	primes := prime(limit)
	var choice []int
	choice = append(choice, 1, 7, 11, 13, 17, 13, 19, 23, 29)
	for {
		k := rand.Intn(int(limit / 30))
		i := choice[rand.Intn(len(choice))]
		c := 30*k + i
		found := true
		for _, v := range primes {
			if c%v == 0 {
				found = false
				break
			}
			if found {
				return c
			}
		}
	}
}
