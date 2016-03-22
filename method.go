package main
import (
	"fmt"
	"math"
)

type Rectangle struct {
	width,height float64	
}

type Circle struct {
	redius float64	
}


//如果method名相同，但接收者不同，那么method就不一样
//方法名area，接收者Rectangle
func (r Rectangle) area() float64 {
	return r.width*r.height	
}
//方法名area，接收者Circle
func (c Circle) area() float64 {
	return c.redius * c.redius * math.Pi	
}


func main(){
	r1 := Rectangle{12,2}
	r2 := Rectangle{9,4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is: ",r1.area())
	fmt.Println("Area of r2 is: ",r2.area())
	fmt.Println("Area of c1 is: ",c1.area())
	fmt.Println("Area of c2 is: ",c2.area())
}	

		
