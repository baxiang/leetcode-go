package main

import "fmt"


// 动态规划
func longestPalindrome2(s string) string {
	if len(s)<2 {
		return s
	}
	dp :=make([][]bool,len(s))
	for i:=range s{
		dp[i] = make([]bool,len(s))
	}
	result := string(s[0])
	for r:=0;r<len(s);r++{
		for l:=0;l<=r;l++{
			if s[l]==s[r]&&(r-l<=2||dp[l+1][r-1]){
				dp[l][r]= true
				if len(s[l:r+1])>len(result){
					result =s[l:r+1]
				}
			}
		}
	}
	return result
}


// 动态规划
func longestPalindrome(s string) string {
    dp :=make([][]bool,len(s))
    for i:=range s{
    	dp[i] = make([]bool,len(s))
	}
    result := ""
    for i :=len(s)-1;i>=0;i--{
    	for j:=i;j<len(s);j++{
				dp[i][j] = s[i]==s[j]&&(j-i<3||dp[i+1][j-1])
				if dp[i][j]&&j-i+1>len(result) {
						result = s[i:j+1]
				}
		}
	}
	return result
}
func longestPalindrome1(s string) string {
	lenth := len(s)

	if lenth <= 1 {
		return s
	}

	dp := make([][]bool, lenth)

	start := 0
	max := 1

	for r := 0;r<lenth;r++ {
		dp[r] = make([]bool, lenth)
		for l:=0;l<r;l++ {
			dp[l][r]= s[l] == s[r] && (r -l <= 2 || dp[l+1][r-1])
			if dp[l][r] {
				curr := r-l+1
				if curr > max {
					max = curr
					start = l
				}
			}
		}
	}
	return s[start:start+max]
}

func main() {
	fmt.Println(longestPalindrome2("cabbac"))
}
