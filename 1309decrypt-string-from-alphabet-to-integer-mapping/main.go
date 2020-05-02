package main

import (
	"fmt"
	"strconv"
)

func freqAlphabets(s string) string {
	 i:=len(s)-1
	 var res =""
     for i>=0{
     	if s[i]=='#'{
     		b:=s[i-2:i]
     		v,_:= strconv.Atoi(b)
     		res= fmt.Sprintf("%s%s",string('a'+v-1),res)
     		i=i-3
		}else {
			v,_:= strconv.Atoi(string(s[i]))
			res= fmt.Sprintf("%s%s",string('a'+v-1),res)
			i--
		}
	 }
	 return res
}

func main() {
	fmt.Println(freqAlphabets("10#11#12"))
	fmt.Println(freqAlphabets("1326#"))

	fmt.Println(freqAlphabets("25#"))
}
