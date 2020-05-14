package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


type Codec struct {

}

func Constructor() Codec {
    return  Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
   queue := make([]*TreeNode,0)
   queue = append(queue,root)
   var res strings.Builder
   res.WriteString("[")
   for len(queue)>0{
   	   l := len(queue)
   	   for i:=0;i<l;i++{
   	   	 if queue[i]!=nil{
			 res.WriteString(fmt.Sprintf("%d,",queue[i].Val))
		 }else {
			 res.WriteString(fmt.Sprintf("null,"))
		 }

	   }
   }
   return res.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	data = data[1:len(data)-1]
	strList  := strings.Split(data,",")
	var root *TreeNode
	for i:=0;i<len(strList)/2;i++{
		v,_ :=strconv.Atoi(strList[i])
		node := &TreeNode{Val: v}
		if strList[2*i+1]!="null"{
			left,_ :=strconv.Atoi(strList[2*i+1])
			node.Left = &TreeNode{Val: left}
		}
		if strList[2*i+2]!="null" {
			right,_ :=strconv.Atoi(strList[2*i+2])
			node.Right = &TreeNode{Val: right}
		}
		if i==0{
			root = node
		}
	}
	return root
}


func main() {
	s := "[1,2,3,null,null,4,5]"

	fmt.Println(s)

	var list []*TreeNode
	list = append(list,nil)
	for _,v:=range list{
		if v==nil{
			fmt.Println("xxxx")
		}
	}
}
