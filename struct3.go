package main 
import "fmt"

type Skills []string

//sturct不仅仅能够将struct作为匿名字段、自定义类型、内置类型都可以作为匿名字段
//而且可以在相应的字段上面进行函数操作
type Human struct {
	name string
	age int
	weight int
}

type Student struct{
	Human  //匿名字段，struct
	Skills //匿名字段，自定义的类型string slice
	int    //内置类型作为匿名字段
	speciality string
}

func main(){
	//初始化学生Jane
	jane := Student{Human:Human{"Jane",35,100},speciality:"Biology"}
   
	fmt.Println("Her name is ",jane.name)
	fmt.Println("Her age is ",jane.age)
	fmt.Println("Her weight is ",jane.weight)
	fmt.Println("Her speciality is ",jane.speciality)

	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ",jane.Skills)
	fmt.Println("She acquired two new ones ")
	jane.Skills = append(jane.Skills,"physice","golang")
	fmt.Println("Her skills now are ",jane.Skills)
	//修改匿名内置类型字段
	jane.int = 3
	fmt.Println("Her preferred number is ",jane.int)
	
}
