package main

import "fmt"
import "time"
/*
The sum of the squares of the first ten natural numbers is,

12 + 22 + ... + 102 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/


func main(){
	starttime := time.Now();
	sum :=0
	sumsquares :=0
	for i:=1; i<=100; i++{
		sum = sum + i
		sumsquares = sumsquares + i*i
	}
	val := sum*sum - sumsquares
	fmt.Println(val)
	fmt.Println(time.Since(starttime));
}