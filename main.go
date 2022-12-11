package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){

	str1 := "goolanggg"
	str2 := "learnnnngoolannnngg"
	fmt.Println(stringCompression(str1))
	fmt.Println(stringCompression(str2))
}



func stringCompression(str string)string{
  count := 1 
	res := ""
	l := len(str)
	for i := 0; i < l; i++ {
		cur := string(str[i])
		next := ""
		if i != l-1 {
			next = string(str[i + 1])
		}
		if strings.Compare(cur, next) == 0 {
			count++
		}else{
			res += cur + strconv.Itoa(count)
			count = 1
		}
	}
return res
}