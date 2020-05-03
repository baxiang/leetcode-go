package main

import "fmt"

func countCharacters(words []string, chars string) int {
	var res int
	m := make(map[byte]int)
	for i :=range chars{
		m[chars[i]]++
	}
	for _,w:=range words{
		if isContainer(w,m) {
			res +=len(w)
		}
	}
	return res
}

func isContainer(w string,m map[byte]int)bool{
	charMap :=make(map[byte]int)
	for i:=range w{
		charMap[w[i]]++
	}

	for k,v :=range charMap{
		if m[k]<v{// 词汇表字母数量小于了当前单词数量 肯定是失败的
			return false
		}
	}
	return true
}

func isContainer1(w string,m map[byte]int)bool{
      charMap :=make(map[byte]int)
      for k,v:=range m{
      	charMap[k]=v
	  }
	for i :=range w{
		if charMap[w[i]]<1{
			return false
		}else {
			charMap[w[i]]--
		}
	}
	return true
}
func main() {
	fmt.Println(countCharacters([]string{"cat","bt","hat","tree"},"atach"))

	fmt.Println(countCharacters([]string{"hello","world","leetcode"},"welldonehoneyr"))

}
