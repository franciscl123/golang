/*
---------------------------------------------------------------------------------------------------------------------------------------------------------------------
author           : Francis Louis
Program Name     : palindromeproduct.go
created date     : 16-Sep-2024
last modified on : 16-Sep-2024
GitHub Repo      : https://github.com/franciscl123/golang.git
Purpose          : https://projecteuler.net/problem=4    (Solution to this problem)  
----------------------------------------------------------------------------------------------------------------------------------------------------------------------
*/

package main

import "fmt"
import  "strconv"

func ispalindrome(n int) bool {
        strval := strconv.Itoa(n)
        if reverse(strval) == strval{
        return true
        }else{
	return false
        }
}

func reverse(s string) string {
    n := len(s)
    runes := make([]rune, n)
    for _, rune := range s {
        n--
        runes[n] = rune
    }
    return string(runes[n:])
}



func main() {
        var x int = 999
        var y int = 999
        var max int = 0
        for i := x; i > 99; i-- {
          for j := y; j > 99; j-- {
             if ispalindrome(i * j){
               if (i * j > max){
                  max = i * j
             }
          }
        }
       }
fmt.Println(max)
}

