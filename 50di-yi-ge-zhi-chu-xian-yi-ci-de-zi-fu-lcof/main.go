package main

func firstUniqChar(s string) byte {


	m := make(map[byte]int)
	keys :=make([]byte,0)
	for i:=range s{
		if _,ok:=m[s[i]];!ok{
			keys = append(keys,s[i])
		}
		m[s[i]]++
	}
	for _,v:=range keys{
		if m[v]==1{
			return v
		}
	}
	var res byte = ' '
	return res
}

func main() {
	
}
