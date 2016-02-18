package main

func main(){
	i := 0
Here:
	println(i)
	i++
	if i <= 10 {
		goto Here
	}
}
