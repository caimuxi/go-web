package main

import "fmt"
/*
func main(){
	i := 10
	switch i {
		case 1:
			fmt.Println("i is equal to 1");
		case 2,3,4:
			fmt.Println("i is equal to 2, 3 or 4");
		case 10:
			fmt.Println("i is equal to 10");
		default:
			fmt.Println("All I know is that i is an integer");
	}
}
*/
//go里面的switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch
//可以使用fallthrough强制执行后面的case代码
func main(){
	integer := 6
	switch integer {
		case 4:
			fmt.Println("The integer was <= 4")
			fallthrough
		case 5:
			fmt.Println("The integer was <= 5")
			fallthrough
		case 6:
			fmt.Println("The integer was <= 6")
			fallthrough
		case 7:
			fmt.Println("The integer was <= 7")
			//fallthrough //默认break
		case 8:
			fmt.Println("The integer was <= 8")
			fallthrough
		default:
			fmt.Println("defaulf case")
	}
}
