/*
---------------------------------------------------------------------------------------------------------------------------------------------------------------------
author           : Francis Louis
Program Name     : multiplessum.go
created date     : 13-Sep-2024
last modified on : 13-Sep-2024
GitHub Repo      : https://github.com/franciscl123/golang.git
Purpose          : https://projecteuler.net/problem=1    (Solution to this problem)  
                 : If we list all the natural numbers below 10  that are multiples of 3 or , 5  we get 3,5,6 and 9. The sum of these multiples is 23 .
                   Find the sum of all the multiples of 3 or 5  below 1000
----------------------------------------------------------------------------------------------------------------------------------------------------------------------
*/

package main

import "fmt"
//import "time"

func main() {
var muls[2] int
var divs[1000] int
var num int
//num = 10
num = 1000
var i int
var j int
var k int
var l int
var sum int
i = 1
j = 0
k = 0
l = 0
sum = 0
muls[0] = 3
muls[1] = 5

for i < num {
 j = 0
 for j < 2{
 if (i % muls[j] == 0){
    divs[k] = i 
    k = k + 1
    break
   }
 j = j + 1
}
i = i + 1
}

for l<=k{
//time.Sleep(1 * time.Second) 
sum = sum + divs[l]
l = l + 1
}
fmt.Println(sum)
}

