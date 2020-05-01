package main

import (
	"fmt"
	"math"
)

func maxNumberOfBalloons(text string) int {
	l :=[26]int{}
	for _,b :=range text{
		l[b-'a']++
	}
	l['l'-'a']= l['l'-'a']/2
	l['o'-'a']= l['o'-'a']/2
	var min = math.MaxInt16
	for _,b :=range "balon"{
		  v :=l[b-'a']
          if v==0{
          	return 0
		  }

		if v<min {
			min = v
		}
	}
	return min
}

func main() {
	fmt.Println(maxNumberOfBalloons("leetcode"))

}
