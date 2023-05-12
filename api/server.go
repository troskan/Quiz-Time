package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Question struct {
	ID      string   `json:"id"`
	Text    string   `json:"text"`
	Choices []string `json:"choices"`
}

var questions = []Question{
	{ID: "1", Text: "What is the capital of France?", Choices: []string{"Paris", "Berlin", "London", "Rome"}},
	// add more questions
}

var answers = map[string]string{
	"1": "Paris",
	// add more answers
}

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a request for questions")
	json.NewEncoder(w).Encode(questions)
}

func SubmitAnswers(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

func GetResults(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

func GetPerformance(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/questions", GetQuestions).Methods("GET")
	router.HandleFunc("/answers", SubmitAnswers).Methods("POST")
	router.HandleFunc("/results/{userID}", GetResults).Methods("GET")
	router.HandleFunc("/performance/{userID}", GetPerformance).Methods("GET")

	http.ListenAndServe(":8080", router)
}
