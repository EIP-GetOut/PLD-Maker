package main

import (
	"fmt"
)

func cutStr(str string, max int) string {
	var result string = ""
	var lenght int = len(str)
	// j define the further position, k is used to find previous space when no '\n' is found in max range.
	var i, j, k int
	var skip bool

	// i define start of substring
	for i = 0; i < lenght; i++ {
		skip = false
		for j = i; skip == false && j < lenght && j < i+max; j++ {
			if str[j] == '\n' {
				result += str[i : j+1]
				i = j
				skip = true
			}
		}
		for k = j; skip == false && k < lenght && k > 0 && k > i; k-- {
			if str[k] == ' ' {
				result += str[i:k] + "\n"
				i = k
				skip = true
			}
		}
		if skip == false {
			if j == lenght {
				result += str[i:lenght]
				break
			} else {
				result += str[i:j+1] + "\n"
				i = j
			}
		}
	}
	return result
}

func main() {
	fmt.Println(cutStr("bonjour je ne sais pas ce que j'écris\nque pense tu de ceci.", 10))
	// fmt.Println(cutStr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", 10))
	//fmt.Println(cutStr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa bonjour je ne sais pas ce que j'écris\nque pense tu de ceci.", 10))
}
