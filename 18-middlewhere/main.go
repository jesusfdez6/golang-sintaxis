package main

import "fmt"

func main(){
	
	result := operacionMidd(sumar)(2,3)
	fmt.Println(result)
	result := operacionMidd(restar)(2,3)
	fmt.Println(result)
}

func operacionMidd(f func(int, int) int) func(int,int) int{
	
	return func(a,b int) int{
	
		fmt.Println("inicio de operacion")
		return f(a,b)
	
	} 

}

func sumar(a,b int) int {

	return a+b
}

func restar(a,b int) int {

	return a-b
}

func multiplicar(a,b int) int {

	return a*b
}

