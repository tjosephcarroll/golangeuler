package main

import "fmt"
import "time"

func main(){
	starttime := time.Now();
	//all the primes under 20
	primes := []int{2,3,5,7,11,13,17,19}
	
	//initialize the output to 1
	out := 1
	var div int
	for i := range primes {
		div = primes[i]
		for div <= 20 {
			div *=primes[i]
		}
		div /= primes[i]
		out *= div
	}
	fmt.Println(out)
	fmt.Println(time.Since(starttime));
}

