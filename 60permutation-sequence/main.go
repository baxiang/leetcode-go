package main

import (
	"fmt"
	"strconv"
	"strings"
)
var mark []int
func getPermutation(n int, k int) string {
     var nums []string
     for i:=1;i<=n;i++{
     	nums = append(nums,strconv.Itoa(i))
	 }
	 path :=make([]string,0)
	 res :=make([]string,0)
	 mark = make([]int,n)
	 dfs(nums,path,k,&res)
	 //fmt.Println(len(res))
	 return res[k-1]
}

func dfs(nums []string,path []string,length int ,res *[]string){
	if len(path)==len(nums) {
		tem := make([]string,len(nums))
		copy(tem,path)
		s := strings.Join(tem,"")
		*res = append(*res,s)
		return
	}
	if len(*res)==length {
		return
	}
	for i :=0;i<len(nums);i++{
		if mark[i]==1 {
			continue
		}
		mark[i]=1
		path = append(path,nums[i])
		dfs(nums,path,length,res)
		path = path[:len(path)-1]
		mark[i]=0
	}
}

func main() {
	fmt.Println(getPermutation(3,3))
	fmt.Println(getPermutation(4,9))
	fmt.Println(getPermutation(9,13531))
}
