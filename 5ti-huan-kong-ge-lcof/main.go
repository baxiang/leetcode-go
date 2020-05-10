package main

import (
	"fmt"
	"strings"
)

func replaceSpace(s string) string {
	var res strings.Builder
    for i:=range s{
    	if s[i]==' '{
			res.WriteString("%20")
		}else {
			res.WriteByte(s[i])
		}
	}
	return res.String()
}


func main() {
	fmt.Println(replaceSpace("We are happy."))
}
