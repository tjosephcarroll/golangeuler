package main
 
import "fmt"
import "time"
import "strconv"
//import "math"

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func checkpalindrome(n int) bool{
	s := strconv.Itoa(n);
	rs := Reverse(s);
	if (s == rs){
		return true;
	}
	return false;	
}

func main(){
	starttime := time.Now();
	max:=0;
	for i:=100; i<1000; i++{
		for j:=100; j<1000; j++{
			num := i*j;
			if ((num>max) && (checkpalindrome(num))){
				max = num;
			}
		}
	}
	fmt.Println(max);
	fmt.Println(time.Since(starttime));
}