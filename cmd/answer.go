/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// answerCmd represents the answer command
var answerCmd = &cobra.Command{
	Use:   "answer",
	Short: "Request user input for question.",
	Long: `Answer will request a user input for question. Only one answer is correct and only one input is allowed.`,
	Run: func(cmd *cobra.Command, args []string) {
var answer string

		fmt.Println("Please enter the number of the answer to send answer.")
		_, err := fmt.Scanln(answer)
		if err != nil {
			fmt.Println("Error reading input:", err)
		}

		  fmt.Println("Your answer is: ", answer)

	},
}

func init() {
	rootCmd.AddCommand(answerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// answerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// answerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
