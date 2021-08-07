package main

import (
	"fmt"
)

var i int


func main() {

	
	for i :=0 ; i <= 10 ;i++  {
		
		fmt.Println(i)
	}


	RUTINA:
		for i<10{

			if i==4{
				i=i+2
				fmt.Println("voy a RUTINA")
				goto RUTINA
			}

			fmt.Println("valor de i: %d\n",i)
			i++


		}
	

}

//break rompe el ciclo for
//continue sigue hasta el bucle
