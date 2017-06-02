package main

import "fmt"
import "time"

func fibonaci(i int)(ret int){
	if i==0 {
		return 0;
	}
	if i==1 {
		return 1;
	}
	
	return fibonaci(i-1) + fibonaci(i-2);
}

func main(){
	starttime := time.Now();
	sum :=0;
	fib :=0;
	for i:=1; i<4000000; i++ {
		fib = fibonaci(i);
		if (fib>4000000){
			fmt.Println(sum);
			fmt.Println(time.Since(starttime));
			return;	
		}
		if ((fib%2)==0){
			sum = sum + fib
		}
	}
}

