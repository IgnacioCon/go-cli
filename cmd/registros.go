package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var registrosCmd = &cobra.Command{
	Use:   "registros",
	Short: "Mostrar el registro de toda la informacion ingresada",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("-- Registros guardados --")

		err := obtenerRegistros()

		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(registrosCmd)
}

func obtenerRegistros() error {
	data, err := ioutil.ReadFile("registros.txt")
	if err != nil {
		return err
	}

	fmt.Println(string(data))

	return nil
}
