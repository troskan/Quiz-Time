/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// answerCmd represents the answer command
var answerCmd = &cobra.Command{
	Use:   "answer",
	Short: "Request user input for question.",
	Long: `Answer will request a user input for question. Only one answer is correct and only one input is allowed.`,
	Run: func(cmd *cobra.Command, args []string) {
		var answer string

		fmt.Println("Please enter the number of the answer to send the answer.")
	
		_, err := fmt.Scanln(&answer)
		if err != nil {
			fmt.Println("Error reading input:", err)
		}
	
		fmt.Println("Your answer is:", answer)
		answerObj := map[string]string{
			"id":     "1",  // We assume the question id is 1 for now
			"answer": answer,
		}
		jsonValue, _ := json.Marshal(answerObj)
		resp, err := http.Post("http://localhost:8080/answer", "application/json", bytes.NewBuffer(jsonValue))

		if err != nil {
				log.Fatalln(err)
		}
		
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
				log.Fatalln(err)
		}
		
		fmt.Printf("Response: %s\n", body)
		defer resp.Body.Close()

		

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
