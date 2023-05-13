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
	CorrectAnswer string `json:"correctAnswer"`
}

type PublicQuestion struct {
	ID      string   `json:"id"`
	Text    string   `json:"text"`
	Choices []string `json:"choices"`
}

type Answer struct {
	ID      string   `json:"id"`
	Answer  string   `json:"answer"`
}

var questions = []Question{
	{ID: "1", Text: "Which company does not produce chainsaws? ", Choices: []string{"Husqvarna", "Partner", "Stihl", "Yamaha"}, CorrectAnswer: "Yamaha"},

	{ID: "2", Text: "Marathon Des Sables is a extreme marathon in the desert, how far is the race distance?", Choices: []string{"150 km", "200 km", "250 km", "300 km"}, CorrectAnswer: "250"},
	
	{ID: "3", Text: "How old is Bill Gates?", Choices: []string{"67", "72", "75", "78"}, CorrectAnswer: "67"},
	// add more questions
}

var answers = map[string]string{
	"1": "Paris",
	"2": "Paris",
	"3": "Paris",
	// add more answers
}


func GetQuestions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a request for questions")
	
	//Hide correct answer
	var publicQuestions []PublicQuestion
	for _, v := range questions {
		publicQuestions = append(publicQuestions, PublicQuestion{ID: v.ID, Text: v.Text, Choices: v.Choices})
	}

	json.NewEncoder(w).Encode(publicQuestions)
}

func SubmitAnswers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a request for answers")

	var answer Answer

	err := json.NewDecoder(r.Body).Decode(&answer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	fmt.Println("Error")
		return
	}
	fmt.Println("Checking if correct answer...")

	for _, v := range questions {
		if answer.ID == v.ID  && answer.Answer == v.CorrectAnswer{
			fmt.Println("Correct answer!")

			fmt.Fprintf(w, "Correct answer!\n")
			
			} else{
				fmt.Fprintf(w, "Wrong answer! Correct answer was: %s\n", v.CorrectAnswer)
	fmt.Println("Wrong answer...")

			}
		
	}
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
	router.HandleFunc("/answer", SubmitAnswers).Methods("POST")
	router.HandleFunc("/results/{userID}", GetResults).Methods("GET")
	router.HandleFunc("/performance/{userID}", GetPerformance).Methods("GET")

	http.ListenAndServe(":8080", router)
}
