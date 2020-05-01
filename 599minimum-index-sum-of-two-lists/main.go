package main

import "fmt"

func findRestaurant(list1 []string, list2 []string) []string {
      m:= make(map[string]int)
      for i,v:=range list1{
      	 m[v] = i
	  }
	  sum := 1<<63-1
	  var res []string
	  for i,v:=range list2{
		  if j,ok:=m[v];ok {
			  if i+j<sum {
				 sum = i+j
				 res =res[:0]
				 res = append(res,v)
			  }else if i+j==sum{
				  res = append(res,v)
			  }
		  }
	  }
	  return res
}

func main() {
	l1 :=[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	l2 :=[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}
	fmt.Println(findRestaurant(l1,l2))
	l11 :=[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	l22 :=[]string{"KFC", "Shogun", "Burger King"}
	fmt.Println(findRestaurant(l11,l22))
}
