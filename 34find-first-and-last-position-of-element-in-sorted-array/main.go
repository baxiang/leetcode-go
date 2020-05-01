package main

import "fmt"

func searchRange(nums []int, target int) []int {
    l :=0
    r :=len(nums)-1
    res :=[]int{-1,-1}
    for l<=r{
    	mid :=l+(r-l)>>1
		if nums[mid]==target{
			res[0]= mid
			res[1] = mid
			i := mid-1
			for i>=0 &&nums[i]==target{
				res[0]=i
				i--
			}
			j := mid+1
			for j<len(nums)&&nums[j]==target  {
				res[1]=j
				j++
			}
			return res
		}else if target>nums[mid] {
			l = mid+1
		}else if target<nums[mid]{
			r = mid-1
		}

    }
    return res
}


func main() {
	fmt.Println(searchRange([]int{5,7,7,8,8,10},8))
}
