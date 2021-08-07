package main

import (
	"fmt"
	"time"
	us "./user"
)


type pepe struct{
	
	us.Usuario
}

func main() {

	User := new(pepe)
	User.AltaUsuario(1,"jesus"time.Now(),true)
	fmt.Println(User)

}

//interface hace definir conductas
