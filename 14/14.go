/*

The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.


*/



package main

import "fmt"
import "time"
//import "io/ioutil"
//import "strconv"
//import "strings"
//import "math"

func main(){
	starttime := time.Now()
	num := 0
	max := 0
	for i:=1; i<1000000; i++{
		val := Collatz(i)
		if val>max{
			max = val
			num = i
		}
	}
	fmt.Println(num)
	fmt.Println(time.Since(starttime))		
}

func Collatz(num int) int {
	//end of the chain
	if (num == 1){
		return 1
	}
	//even
	if (num%2 == 0){
		val := num/2
		return (1 + Collatz(val))
	}
	
	//odd
	val := num*3 +1
	return (1 + Collatz(val))
}