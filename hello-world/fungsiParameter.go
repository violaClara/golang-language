package main

import {
	"fmt"
	"strings"
}

func main(){
	var data=[]string{"wick","jason","ethan"}
	var dataContainsO= filter(data, func(each string) bool{
		return strings.Contains(each,"o")
	})
	var dataLength5= filter(data, func(each string)bool{
		return len(each)==5
	})

	fmt
}



func filter(data []string, callback func (string) bool)[]string{
	var result[]string
	for _,each:= range data{
		if filtered := callback(each);filtered{
			result=append(result,each)
		}

	}
	return result
}