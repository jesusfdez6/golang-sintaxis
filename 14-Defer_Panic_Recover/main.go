package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
   
)

func main(){

	archivo,err := os.Open("./archivo.txt")

	defer archivo.Close()

	if err != nil {
		
		fmt.println("hubo un error")
		os.exit(1)
	}
	ejemploPanic()

}
func ejemploPanic(){
	defer func(){
		reco := recover()
		if reco !=nil{
			log.Fatalf("error panic \n %v",reco)		
	}()
	a := 1
	if a==1 {
		panic("se encontró el valor de 1")
	}
}


//DEfer hace ese codigo sin importar que haya ocurrido un error
//Panic
//recover se va a ejecutar cuando haya un panic