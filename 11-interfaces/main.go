package main

import "fmt"

type serVivo interface{
	estaVivo() bool
}

type humano interface{
	
	respirar()
	pensar()
	comer()
	sexo() string
	estaVivo() bool
}

type animal interface{

	respirar()
	comer()
	Escarnivoro() bool
	estaVivo() bool
}

type vegetal interface{

	ClasificacionVegetal() string

}

type perro struct{

	respirando bool
	comiendo bool
	carnivoro bool
	vivo bool
}

type hombre struct{

	edad int
	altura float32
	peso float32
	respirando bool
	pensando bool
	comiendo bool
	vivo bool
}

type mujer struct{

	edad int
	altura float32
	peso float32
	respirando bool
	pensando bool
	comiendo bool
	vivo bool
}

func (h *hombre) respirar(){   h.respirando=true}
func (h *hombre) comer(){   h.comiendo=true}
func (h *hombre) pensar(){   h.pensando=true}
func (h *hombre) sexo() string{   return "hombre"}
func (h *hombre) estaVivo() bool{   return h.vivo}


func (m *mujer) respirar(){   m.respirando=true}
func (m *mujer) comer(){   m.comiendo=true}
func (m *mujer) pensar(){   m.pensando=true}
func (m *mujer) sexo() string {   return "mujer"}
func (m *mujer) estaVivo() bool{   return m.vivo}


func (p *perro) comer(){   p.comiendo=true}
func (p *perro) respirar(){   p.respirando=true}
func (p *perro) Escarnivoro() bool {   return p.carnivoro}
func (p *perro) estaVivo() bool{   return p.vivo}


func humanoRespirando(hu humano){

	hu.respirar()
	fmt.Printf("soy un/a %s, y ya estoy respirando \n",hu.sexo())
}

func animalRespirando(a animal){

	a.respirar()
	fmt.Printf("soy un animal y estoy respirando")
}

func animalesCarnivoros(an animal) int {

	if an.Escarnivoro()==true{
		return 1
	}
	return 0
}

func estoyVivo(s serVivo) bool{
	 return s.estaVivo()
}


func main() {

	hombre := new(hombre)
	hombre.vivo=false
	humanoRespirando(hombre)
	mujer := new(mujer)
	humanoRespirando(mujer)
	totalcarnivoros :=0
	Dogo := new(perro)
//	animalRespirando(Dogo)
	Dogo.carnivoro =true
	totalcarnivoros += animalesCarnivoros(Dogo)
	Dogo1 := new(perro)
	Dogo1.carnivoro =true
	totalcarnivoros += animalesCarnivoros(Dogo1)
	fmt.Println(totalcarnivoros)
	fmt.Println(estoyVivo(hombre))





}
