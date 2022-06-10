/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long:  `This is serve command`,
	Run: func(cmd *cobra.Command, args []string) {
		inputName, _ := cmd.Flags().GetString("name")
		inputAge, _ := cmd.Flags().GetString("age")
		fmt.Println(inputName + inputAge)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.PersistentFlags().String("name", "dummy", "Enter User Name")
	serveCmd.PersistentFlags().String("age", "0", "Enter User Age")
}
