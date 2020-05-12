package main

import "fmt"

func addStrings(num1 string, num2 string) string {
	l1 := len(num1)-1
	l2 := len(num2)-1
	carry :=0
	str :=""
	for l1>=0||l2>=0{
		n1 :=0
		if l1>=0{
			n1 = int(num1[l1]-'0')
			l1--
		}
		n2 :=0
		if l2>=0{
			n2 = int(num2[l2]-'0')
			l2--
		}
		fmt.Println(n2,n1)
		r := n1+n2+carry
		carry= r/10
        str =fmt.Sprintf("%d%s",r%10,str)
	}
	if carry>0{
		str =fmt.Sprintf("1%s",str)
	}
	return str
}

func main() {
	fmt.Println(addStrings("900","211"))
}
