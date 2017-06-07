package main

import "fmt"
import "time"
//import "strconv"
//import "math"

/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a2 + b2 = c2
For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

func main(){
	starttime := time.Now();
	
	/*
	   a = m2 - n2
       b = 2 * m * n
       c  = m2 + n2
	   because,
       a2 = m4 + n4 – 2 * m2 * n2
       b2 = 4 * m2 * n2
       c2 = m4 + n4 + 2* m2 * n2
	*/
	
	a := 0
	b := 0
	c := 0
	
	m := 2
	
	for c < 1000 {
		for n:=1; n < m; n++{
			
			a = m*m - n*n
			b = 2*m*n
			c = m*m + n*n	
			if (c > 1000){
				break
			}
			if (a+b+c == 1000){
				fmt.Println(a,b,c,a*b*c)
			}
			
			
		}
		m++
	}
	
	fmt.Println(time.Since(starttime));	
}