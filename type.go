package main
import "fmt"

//适用于通用接口
//通过type将函数定义成变量
type testInt func(int) bool //声明一个函数类型

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false	
	}	
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true	
	}	
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value){
			result = append(result,value)	
		}	
	}
	return result
}

func main(){
	slice := []int {1,2,3,4,5,6,7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)
	fmt.Println("Old elements of slice are: ",odd) //函数当做值来传递
	even := filter(slice,isEven)
	fmt.Println("Even elements of slice are: ",even)

}
