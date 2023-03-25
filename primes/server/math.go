package main

import "log"

func calcPrimes(n int32, next func(i int32)) {
	var k int32 = 2
	for n > 1 {
		log.Printf("n = %d", n)
		if n%k == 0 {
			// if k evenly divides into N
			// this is a factor
			next(k)
			log.Printf("next k=%d", n)
			// divide n by k so that we have the rest of the number left
			n = n / k
			log.Printf("next n=%d", n)
		} else {
			k = k + 1
		}
	}
}
