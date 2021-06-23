package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ignaciocon/exam-cli/models"
	"github.com/spf13/cobra"
)

// agregarPersonaCmd represents the agregarPersona command
var agregarPersonaCmd = &cobra.Command{
	Use:   "agregarPersona",
	Short: "Crear un nuevo registro de persona",
	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Printf("--Ingrese los siguientes valores--\n\n")

		var nombre string

		fmt.Print("Nombre Completo: ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			nombre = scanner.Text()
		}

		var edad int
		fmt.Print("Edad: ")
		_, err := fmt.Scanln(&edad)

		if err != nil {
			return fmt.Errorf(err.Error())
		}

		var sexo string
		fmt.Print("Sexo (H hombre, M mujer): ")
		_, err = fmt.Scanln(&sexo)

		if err != nil {
			return fmt.Errorf(err.Error())
		}

		var peso float64
		fmt.Print("Peso (en kg): ")
		_, err = fmt.Scanln(&peso)

		if err != nil {
			return fmt.Errorf(err.Error())
		}
		var altura float64
		fmt.Print("Altura (en cm): ")
		fmt.Scanln(&altura)

		if err != nil {
			return fmt.Errorf(err.Error())
		}

		persona := models.NuevaPersona(nombre, edad, sexo, peso, altura)

		fmt.Printf("--------------------------------\n\n")

		//Calcular el IMC de la persona
		IMC := persona.CalcularIMC()

		estado := "La persona tiene un IMC que indica que tiene "

		if IMC == 1 {
			fmt.Println(estado + models.SobrePeso + ".")
		} else if IMC == 0 {
			fmt.Println(estado + models.PesoNormal + ".")
		} else if IMC == -1 {
			fmt.Println(estado + models.FaltaDePeso + ".")
		}

		//Comprobar si la persona es mayor de edad
		if persona.EsMayorDeEdad() {
			fmt.Printf("La persona es mayor de edad.\n\n")
		} else {
			fmt.Printf("La persona no es mayor de edad.\n\n")
		}

		fmt.Println(persona.String())
		fmt.Println("--------------------------------")

		mensaje, err := guardarPersona(persona)

		if err != nil {
			return err
		}

		fmt.Println(mensaje)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(agregarPersonaCmd)
}

func guardarPersona(persona models.Persona) (string, error) {
	file, err := os.OpenFile("registros.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return "Error al guardar.", err
	}

	defer file.Close()

	_, err = file.WriteString(persona.String())

	if err != nil {
		return "Error al guardar.", err
	}

	return "Persona guardada con exito!", nil

}
