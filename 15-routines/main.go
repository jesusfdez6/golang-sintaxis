package main

import (
	"fmt"
	"strings"
	"time"
)

func main(){

	//lo va hacer de manera asincrona por el comando go
	go miNombreLentooo("jesus")

}

func miNombreLentooo(nombre string){
	letras := strings.Split(nombre,"")
	for _,letra := range letras{

		time.Sleep(1000*time.Millisecond)
		fmt.Println(letra)
	}

}