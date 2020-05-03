package main


func maximum(a int, b int) int {
	 k := ((a - b) >> 63) & 1
	 return b * k + a * (k ^ 1)
}
func main() {
	
}
