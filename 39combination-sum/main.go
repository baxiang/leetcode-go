package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
  res := make([][]int,0)
  path := make([]int,0)
  dfs(candidates,target,path,&res,0)
  return res
}

func dfs(candidates []int, target int,path[]int,res*[][]int,idx int){
	if target==0 {
		*res=append(*res,append([]int{},path...))
	}
	for i:=idx;i<len(candidates);i++{
		if target-candidates[i]>=0 {
			path = append(path,candidates[i])
			dfs(candidates,target-candidates[i],path,res,i)
			path = path[:len(path)-1]
		}
	}
}

func main() {
	fmt.Println(combinationSum([]int{2,3,6,7},7))
}
