package main

import "fmt"

func judgeCircle1(moves string) bool {
     m :=make(map[byte]int)
     for i:=range moves{
     	m[moves[i]]++
	 }
	if m['U']==m['D']&&m['R']==m['L'] {
		return true
	}
	return false
}

func judgeCircle(moves string) bool {
	m :=make([]int,4)
	for _,v:=range moves{
		if v=='U'{
			m[0]++
		}else if v=='D'{
			m[1]++
		}else if v=='R'{
			m[2]++
		}else {
			m[3]++
		}
	}
	if m[0]==m[1]&&m[2]==m[3] {
		return true
	}
	return false
}



func main() {
	fmt.Println(judgeCircle("UD"))
	fmt.Println(judgeCircle("LL"))
}
