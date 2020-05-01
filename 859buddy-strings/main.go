package main

import "fmt"

func buddyStrings(A string, B string) bool {
	if len(A)!=len(B) {
		return false
	}

	count := 0
	w := A[0]
	isSameStr := true
	for i :=0;i<len(A)-1;i++{
		if A[i]==B[i] {
			if w !=B[i]&&isSameStr  {
				isSameStr = false
			}
			continue
		}
		if A[i]==B[i+1]&&A[i+1]==B[i] {
			count++
		}
		if count>2 {
			return false
		}
	}
	if isSameStr {
		return true
	}
	return count==1
}

func main() {
	fmt.Println(buddyStrings("ab","ba"))
	fmt.Println(buddyStrings("ab","ab"))
	fmt.Println(buddyStrings("aa","aa"))
	fmt.Println(buddyStrings("aaaaaaabc","aaaaaaacb"))
}
