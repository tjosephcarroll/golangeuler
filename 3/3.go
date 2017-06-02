package main

import "fmt"
import "time"
import "math"

func primefactors(n int) {

	if (n%2 == 0){
		fmt.Println(2);
	}
	for i:=3; i<= int(math.Sqrt(float64(n))); i++ {
		for (n%i == 0) {
			fmt.Println(i);
			n = n/i;
		}
	}
	if (n > 2){
		fmt.Println(n)
	}
}

func main(){
	starttime := time.Now();
	primefactors(600851475143);
	fmt.Println(time.Since(starttime));
}

