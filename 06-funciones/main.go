package main

import (
	"fmt"
)
 
func main() {


//valor :=uno(5)
//valor1,valor2 :=dos(5)
valor :=Calculo(5,20,21)

fmt.Println(valor)
//fmt.Println(valor1)
//fmt.Println(valor2)

}


func uno(numero int) (z int) {
	z = numero*2
	return 
}

func dos(numero int) (int, bool) {
	
	return numero*2,true
}

//no se sabe el numero de parametros
func Calculo(numero ...int) (total int){
	
	total=0
	
	for _,num:= range numero{
		total = total + num
	}
	return 
}



