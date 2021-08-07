package main

import "fmt"

//GLOBAL
var numero int =10
//var texto string // texto
var status bool // booleana
//var decimal float32 // float



func main(){
	
	//var numero2,numero4,numero5 int // entera pero no es global
//	numero3 := 4 
//	numero3 = 5
	numero6,numero7 := 2,5
	numero :=5

	var moneda int64 =0
	texto = strconv.Itoa(int(moneda))
	fmt.Println(numero6)	
	fmt.Println(numero7)
	fmt.Println(status )	
	fmt.Println(numero)	

	mostrarStatus()


}

func MostrarStatus(){
	fmt.Println(numero)
}



// Con la letra mayuscula en la primera letra de la palabra hace que sea publica y en minuscula entonces es privada