package main

import "fmt"

func uniqueMorseRepresentations(words []string) int {
	s :=[26]string {".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
   m := make(map[string]struct{})
   for _,word :=range words {
   	key :=""
   	for _,v :=range word{
   		key=key+s[v-'a']
	}
	   if _,ok:=m[key];!ok {
		  m[key]= struct{}{}
	   }
   }
	return len(m)
}

func main() {
  fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}
