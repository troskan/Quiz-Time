/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// questionCmd represents the question command
var questionCmd = &cobra.Command{
	Use:   "question",
	Short: "A brief description of your command",
	Long: `Will display question`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("http://localhost:8080/questions")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		var questions []Question
		if err := json.NewDecoder(resp.Body).Decode(&questions); err != nil {
			log.Fatal(err)
		}

		for _, question := range questions {
			fmt.Printf("Question: %s\n", question.Text)
			for i, choice := range question.Choices {
				fmt.Printf("#%d: %s\n", i+1, choice)
			}
		}
	},
}

type Question struct {
	ID      string   `json:"id"`
	Text    string   `json:"text"`
	Choices []string `json:"choices"`
}

func init() {
	rootCmd.AddCommand(questionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// questionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// questionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
