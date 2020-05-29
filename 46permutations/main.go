package main

import "fmt"

var mark []int
var n int
func permute(nums []int) [][]int {
    n =len(nums)
    res:= make([][]int,0)
    if n<=0{
    	return res
	}
	mark = make([]int,n)
	path :=make([]int,0)
	dfs(nums,path,&res)
	return res
}

func dfs(nums[]int,path[]int,res*[][]int)  {
	if len(path)==n {
		tmp := make([]int,n)
		copy(tmp,path)
		*res =append(*res,tmp)
		return
	}
	for i :=0;i<n;i++{
		if mark[i]==1 {//减掉枝叶
			continue
		}
		path = append(path,nums[i])
		mark[i]=1
		dfs(nums,path,res)
		path = path[:len(path)-1]
		mark[i]=0
	}
}

func main() {
	fmt.Println(permute([]int{1,2}))
}
