package main

import "fmt"
import "time"

func main(){
	starttime := time.Now();
	sum := 0;
	for a:=0; a<1000; a++ {
		if (((a%3)==0)||((a%5)==0)){
			sum = sum + a;
		}
	}
	
	fmt.Println(sum);
	fmt.Println(time.Since(starttime));
}