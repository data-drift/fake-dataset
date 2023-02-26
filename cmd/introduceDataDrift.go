/*
Copyright © 2023 NAME HERE sammy.teillet@gmail.com

*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"

	"github.com/spf13/cobra"
)

// introduceDataDriftCmd represents the introduceDataDrift command
var introduceDataDriftCmd = &cobra.Command{
	Use:   "introduceDataDrift",
	Short: "Introduce data drift in a table",
	Long:  `Given a csv file, this command will introduce a random modification in the historical data`,
	Run: func(cmd *cobra.Command, args []string) {
		filepath := args[0]
		newFilePath := filepath + ".new"
		csvFile, err := os.Open(filepath)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Successfully Opened CSV file")
		defer csvFile.Close()
		csvLines, err := csv.NewReader(csvFile).ReadAll()
		fmt.Printf("len(csvLines): %d \n", len(csvLines))

		if err != nil {
			fmt.Println(err)
		}
		randomIndex := rand.Intn(len(csvLines)-1) + 1
		fmt.Printf("randomIndex: %d \n", randomIndex)
		newCsvLines := append(csvLines[:randomIndex], csvLines[randomIndex+1:]...)

		newCsvFile, err := os.Create(newFilePath)
		csvwriter := csv.NewWriter(newCsvFile)

		_ = csvwriter.WriteAll(newCsvLines)

		csvwriter.Flush()
		csvFile.Close()

		_ = os.Remove(filepath)
		_ = os.Rename(newFilePath, filepath)
		fmt.Println("Successfully removed a line in CSV file.")

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
