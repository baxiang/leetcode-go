package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	str :="1"
    for i:=2;i<=n;i++{
    	var tmp strings.Builder
    	preByte := str[0]
    	count := 1
    	for j:=1;j<len(str);j++{
    		if str[j]==preByte{
    			count++
			}else {
				tmp.WriteString(strconv.Itoa(count))
				tmp.WriteByte(preByte)
				preByte =str[j]
				count=1
			}
		}
		tmp.WriteString(strconv.Itoa(count))
		tmp.WriteByte(preByte)
    	str = tmp.String()
	}
	return str
}

func main() {
	fmt.Println(countAndSay(4))
}
