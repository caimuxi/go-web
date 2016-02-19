package main

import "fmt"

//一个返回值
/*
func max(a,b int) int {
	if a > b {
		return a	
	}	
	return b
}

func main(){
	x := 3
	y := 4
	z := 5

	max_xy := max(x, y)
	max_xz := max(x, z)

	fmt.Printf("max(%d, %d) = %d\n", x, y, max_xy)
	fmt.Printf("max(%d, %d) = %d\n", x, z, max_xz)
	fmt.Printf("max(%d, %d) = %d\n", y, z, max(y,z)) //直接调用max函数
}
*/

//多个返回值(未命名)
/*
func SumAndProduct(A,B int)(int,int) { //第二个括号里的int是返回值类型
	return A+B,A*B
}
*/
//多个返回值（命名）
func SumAndProduct(A,B int)(add int, Multiplied int) {
	add = A + B
	Multiplied = A * B
	return 
}

func main(){
	x := 3
	y := 4

	xPLUSy, xTIMESy := SumAndProduct(x, y)

	fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
	fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)
}
