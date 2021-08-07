package main

import (
	"fmt"
)


func main() {

paises := make(map[string]string)
paises["Mexico"]="D.F"
paises["Argentina"]="Buenos aires"

	campeonato := map[string]int{
		"barcelona" : 39,
		"Real madrid" : 38,
		"Atletico de madrid" : 35,
		"Valencia" : 30}
	

	for equipo, puntaje :=range campeonato {

		fmt.Printf("El equipo %s, tiene un puntaje de %d" equipo,puntaje)
	}

	



}