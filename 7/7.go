package main

import "fmt"
import "time"
import "math"
/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
What is the 10 001st prime number?
*/
func makeRange(min, max int) []bool{
	a := make([]bool, max-min+1)
    for i := range a {
        a[i] = true
    }
    return a
}


func main(){
	starttime := time.Now();
	
	//Sieve of Eratosthenes 
	n := 1000000
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
	
	//count
	count :=0
	for h := range a {
		if (a[h]){
			count++
			if (count == 10001){
				fmt.Println(h)
			}
		}
		
	}
	fmt.Println(count)
	fmt.Println(time.Since(starttime));	
}