/*
Copyright © 2023 NAME HERE sammy.teillet@gmail.com

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// introduceDataDriftCmd represents the introduceDataDrift command
var introduceDataDriftCmd = &cobra.Command{
	Use:   "introduceDataDrift",
	Short: "Introduce data drift in a table",
	Long:  `Given a csv file, this command will introduce a random modification in the historical data`,
	Run: func(cmd *cobra.Command, args []string) {
		filepath := args[0]
		fmt.Println("introduceDataDrift called", filepath)
	},
}

func init() {
	rootCmd.AddCommand(introduceDataDriftCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// introduceDataDriftCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// introduceDataDriftCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
