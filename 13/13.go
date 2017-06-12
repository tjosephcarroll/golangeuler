/*

Work out the first ten digits of the sum of the following one-hundred 50-digit numbers.

*/
package main

import "fmt"
import "time"
import "io/ioutil"
import "strconv"
import "strings"
//import "math"

func main(){
	starttime := time.Now()
	
	//from 11
	fBuffer, error := ioutil.ReadFile("13.txt")
	if error != nil {
		return
    }
    fString := string(fBuffer)
	lines := strings.Split(fString,"\n",)
    
    result := ""
    carryout, digit := 0, 0
    for i := 49; i >= 0; i-- {
        sum := carryout
        for j := range lines {
            num, _ := strconv.Atoi(string(lines[j][i]))
            sum += num
        }
        digit = sum % 10
        carryout = sum / 10
        result = strconv.Itoa(digit) + result
    }
    result = strconv.Itoa(carryout) + result
    println(result[0:10])
	fmt.Println(time.Since(starttime))		
}