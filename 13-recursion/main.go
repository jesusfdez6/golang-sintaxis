package main

import (
	"fmt"

)

func main(){

}

func exponencia(numero int){

	if numero >100{
		return
	}

	fmt.println(numero)
	exponencia(numero *2)
}