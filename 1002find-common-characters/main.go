package main

import "fmt"

func commonChars(A []string) []string {
     m :=make(map[byte]int,0)
     for idx,v:=range A{
     	if idx==0{
			for i:=range v{
				m[v[i]]++
			}
		}else {
			n :=make(map[byte]int,0)
			for i:=range v{
				n[v[i]]++
			}
			for k,v:=range m{
				 if r,ok:=n[k];ok{
					m[k]=min(v,r)
				}else {
					delete(m,k)
				 }
			}
		}
	 }
	 var res []string
	 for k,v:=range m{
	 	for i:=0;i<v;i++{
	 		res =append(res,string(k))
		}
	 }
	 return res
}

func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}

func main() {
	fmt.Println(commonChars([]string{"bella","label","roller"}))
	fmt.Println(commonChars([]string{"cool","lock","cook"}))
}
