package models

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

const SobrePeso = "Sobrepeso"
const PesoNormal = "Peso normal"
const FaltaDePeso = "Falta de peso"

//Persona - estructura con caracteristicas de una persona.
type Persona struct {
	nombre      string
	edad        int
	nss         string
	sexo        string
	peso        float64
	altura      float64
	fechaCreada time.Time
}

//NuevaPersona - Constructor para crear nueva Persona con parametros
func NuevaPersona(nombre string, edad int, sexo string, peso float64, altura float64) Persona {
	persona := Persona{}

	persona.setNombre(nombre)
	persona.setEdad(edad)
	persona.nss = generarNSS()
	persona.setSexo(sexo)
	persona.setPeso(peso)
	persona.setAltura(altura)
	persona.fechaCreada = time.Now()

	return persona
}

func (p *Persona) setNombre(nombre string) {
	p.nombre = nombre
}

func (p *Persona) setEdad(edad int) {
	if edad < 0 {
		p.edad = 0
		return
	}
	p.edad = edad
}

func (p *Persona) setSexo(sexo string) {
	if sexo == "H" || sexo == "M" {
		p.sexo = sexo
	} else {
		p.sexo = "H"
	}
}

func (p *Persona) setPeso(peso float64) {
	if peso < 0 {
		p.peso = 0
		return
	}
	p.peso = peso
}

func (p *Persona) setAltura(altura float64) {
	if altura < 0 {
		p.altura = 0
		return
	}
	p.altura = altura
}

//CalcularIMC para persona
func (p *Persona) CalcularIMC() int {
	//convertir altura de cm a metros
	estaturaEnMetros := p.altura / 100

	//calcular IMC
	IMC := p.peso / (estaturaEnMetros * estaturaEnMetros)

	var resultado int

	//Comprobar si la persona es hombre o mujer
	switch p.comprobarSexo("H") {

	case true: //es Hombre
		if IMC >= 25 {
			resultado = 1
		} else if IMC >= 20 && IMC < 25 {
			resultado = 0
		} else if IMC < 20 {
			resultado = -1
		}

	case false: //es Mujer
		if IMC >= 24 {
			resultado = 1
		} else if IMC >= 19 && IMC < 24 {
			resultado = 0
		} else if IMC < 19 {
			resultado = -1
		}
	}

	return resultado
}

//Regresa
//true si la persona es mayor de 18 años
//false si la persona no es mayor de 18 años
func (p *Persona) EsMayorDeEdad() bool {
	return p.edad >= 18
}

//Regresa verdadero si el sexo ingresado es correcto
//falso si es incorrecto
func (p *Persona) comprobarSexo(sexo string) bool {
	return p.sexo == sexo
}

func (p *Persona) String() string {
	sexo := "Hombre"

	if p.comprobarSexo("M") {
		sexo = "Mujer"
	}

	return fmt.Sprintf("Nombre: %s\nEdad: %d\nNSS: %s\nSexo: %s\nPeso: %.1fkg\nAltura: %.1fcm\nFecha Creada: %s\n\n", p.nombre, p.edad, p.nss, sexo, p.peso, p.altura, p.fechaCreada.Format(time.RFC822))
}

func generarNSS() string {

	//Crear unn slice de bytes de 4 numeros
	numeros := make([]byte, 4)

	//leer 4 numeros al azar al slice numeros
	_, err := rand.Read(numeros)

	if err != nil {
		fmt.Println("Error reading numbers: ", err.Error())
	}

	//Convertir los numeros a su valor hexadecimal
	nss := hex.EncodeToString(numeros)
	return strings.ToUpper(nss)
}
