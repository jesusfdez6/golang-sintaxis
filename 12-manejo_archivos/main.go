package main

import (
	"fmt"
	"os"
	"bufio"
	"io/ioutil"
)

func main(){

	leoArchivo()
	leoArchivo2()
	graboArchivo()
	graboArchivo2()

}

func leoArchivo(){

	archivo, err := ioutil.ReadFile("./archivo.txt")
	
	if err != nil {
		
		fmt.println("hubo un error")
	}else{

		fmt.println(string(archivo))
	}

}

func leoArchivo2(){

	archivo,err := os.Open("./archivo.txt")

	if err != nil {
		
		fmt.println("hubo un error")
	}else{

		scanner :=	bufio.newScanner(archivo)
		for scanner.Scan(){
			registro := scanner.Text()
			fmt.Printf("Linea >"+registro)
		}
	}

	archivo.Close()
}

func graboArchivo(){
	archivo, err := os.Create("./nuevoArchivo.txt")

	if err != nil{
		fmt.println("Hubo un error")
		return
	}

	fmt.Fprintln(archivo,"esta es una linea nueva")
	archivo.close()
}

func graboArchivo2(){

	filename := "./nuevoArchivo.txt"
	if Append(filename,"\n Esta es una segunda linea") == false {
		fmt.println("Hubo un error")
		return
	}
}

func Append(archivo string, texto string) bool{

	archivo,err := os.OpenFile(archivo, os.O_WRONLY | os.O_APPEND,0644)
	if err != nil{
		return false
	}
  // _ no nos interesa el valor de esa variable que es el resultado de la escritura del arxhivo
	_, err  = archivo.WriteString(texto)

	if err != nil{
		return false
	}

	return true
}