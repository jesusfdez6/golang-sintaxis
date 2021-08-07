package main

import (
	"fmt"
)
 
var Calculo func(int,int) int =func (num1 int,num2 int) int {
	return num1+num2
}

func main() {
fmt.Println(Calculo(1,2))
operaciones()
tabla :=2
miTabla := Tabla(tabla);
for i :=1;i<10;i++{

	fmt.Println(miTabla())

}
}

func operaciones()  {
	resultado := func() int{
		var a int = 23
		var b int = 13
		return a+b
	}
	fmt.Println(resultado())
}


func Tabla(valor int) func() int {
	numero := valor 
	secuencia :=0

	return func() int{
		secuencia +=1
		return numero*secuencia
	} 
}









