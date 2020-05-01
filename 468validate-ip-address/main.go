package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func checkIPv4(IP string)bool{
	str := strings.Split(IP,".")
	if len(str)!=4{
		return false
	}
	for _,s:=range str{
		if len(s)==0||len(s)>1&&s[0]=='0' {
			return false
		}
		if len(s)>0&&!unicode.IsDigit(rune(s[0])){
			return false
		}
		r,err := strconv.Atoi(s)
		if err!=nil {
			return false
		}
		if r <0||r>255{
			return false
		}
	}
	return true
}
func checkIPv6(IP string)bool{
	str := strings.Split(IP,":")
	if len(str)!=8{
		return false
	}
	for _,s :=range str{
		if len(s) <= 0 || len(s) > 4 {
			return false
		}
		for i := 0; i < len(s); i++ {
			if s[i] >= '0' && s[i] <= '9' {
				continue
			}
			if s[i] >= 'A' && s[i] <= 'F' {
				continue
			}
			if s[i] >= 'a' && s[i] <= 'f' {
				continue
			}
			return false
		}
	}
	return true
}

func validIPAddress(IP string) string {
	if checkIPv4(IP) {
		return "IPv4"
	}
	if checkIPv6(IP) {
		return "IPv6"
	}
	return "Neither"
}


func main() {
	//fmt.Println(validIPAddress("172.16.254.1"))
	//fmt.Println(validIPAddress("172.16.254.01"))
	//fmt.Println(validIPAddress("12..33.4"))
	fmt.Println(validIPAddress("192.0.0.1"))
	fmt.Println(validIPAddress("15.16.-0.1"))
	//fmt.Println(validIPAddress("2001:0db8:85a3:0000:0000:8a2e:0370:7334"))
	//fmt.Println(validIPAddress("2001:0db8:85a3::8A2E:0370:7334"))
}
