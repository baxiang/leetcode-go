package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := candies[0]
	for i := 1; i < len(candies); i++ {
		if max < candies[i] {
			max = candies[i]
		}
	}
	res := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		// 当前拥有的糖果数加上额外的糖果数大于等于最大值就可以满足
		if candies[i]+extraCandies >= max {
			res[i] = true
		}
	}
	return res
}

func main() {
	fmt.Println(kidsWithCandies([]int{2,3,5,1,3},3))
    fmt.Println(kidsWithCandies([]int{4,2,1,1,},1))
}
