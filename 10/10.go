package main

import "fmt"
import "time"
//import "strconv"
import "math"

/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/


func main(){
	starttime := time.Now();
	
	//borrowing from 7
	//Sieve of Eratosthenes
	
	n := 2000000
	a := makeRange(1, n)
	a[0] = false
	for i:=2; i<=int(math.Sqrt(float64(n))); i++{
		j:=(i*i)
		k:= 0
		for j<=n{
			a[j-1] = false;
			k = k + 1
			j = i*i + k*i
		}
	}
	
	//sum
	sum :=0
	for h := range a {
		if (a[h]){
			//shift up 1
			sum = sum + (h+1)
		}
	}
	fmt.Println(sum)
	
	
	fmt.Println(time.Since(starttime));	
}

func makeRange(min, max int) []bool{
	a := make([]bool, max-min+1)
    for i := range a {
        a[i] = true
    }
    return a
}
