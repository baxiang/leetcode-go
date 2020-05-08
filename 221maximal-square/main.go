package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	maxSide :=0
	if len(matrix)==0 ||len(matrix[0])==0{
		return maxSide
	}
	for i:=0;i<len(matrix);i++{
		for j:=0;j<len(matrix[i]);j++{
			if matrix[i][j]=='1'{
				maxSide = max(maxSide,1)
				//以i点为左上角 做大的边长
				currMaxSide :=min(len(matrix)-i,len(matrix[i])-j)
				for k:=1;k<currMaxSide;k++{
					flag :=true
					// 以i点为左上角 向右下方的斜角位置点
					if matrix[i+k][j+k]=='0'{
						break
					}
					for  m := 0; m < k; m++ {
						if matrix[i + k][j + m] == '0' || matrix[i + m][j + k] == '0' {
							flag = false
							break
						}
					}
					if flag{
						maxSide = max(maxSide,k+1)
					}else {
						break
					}
				}
			}
		}
	}
	return maxSide*maxSide
}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}

func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}
func main() {
	//g :=[][]byte{
	//	{'1','0','1','0','0'},
	//	{'1','0','1','1','1'},
	//	{'1','1','1','1','1'},
	//	{'1','0','0','1','0'},
	//}
	//fmt.Println(maximalSquare(g))
	//g1 :=[][]byte{
	//	{'1','0','1','1','1'},
	//	{'0','1','0','1','0'},
	//	{'1','1','0','1','1'},
	//	{'1','1','0','1','1'},
	//	{'0','1','1','1','1'},
	//}
	//fmt.Println(maximalSquare(g1))


	g2 :=[][]byte{
		{'1','0','1','1','0','1'},
		{'1','1','1','1','1','1'},
		{'0','1','1','0','1','1'},
		{'1','1','1','0','1','0'},
		{'0','1','1','1','1','1'},
		{'1','1','0','1','1','1'},
	}
	fmt.Println(maximalSquare(g2))
}
