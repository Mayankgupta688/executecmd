/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// sampleCmd represents the sample command
var sampleCmd = &cobra.Command{
	Use:   "sample",
	Short: "A brief description of your command",
	Long:  `This is Sample Command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sample called")
	},
}

func init() {
	rootCmd.AddCommand(sampleCmd)
}
