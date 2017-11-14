package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Seet the random generator once the application is started
	rand.Seed(time.Now().UnixNano())
	//Set of requests
	request := [5]string{
		"People say I look like both my mother and father.",
		"Father was a teacher.",
		"I was my father’s favourite.",
		"I’m looking forward to the weekend.",
		"My grandfather was French!"}
	//Lop the request
	for i, req := range request {
		fmt.Println(i+1, "\nRequest:", req)
		//Get the response
		fmt.Println("Response:", ElizaResponse(req), "\n")
	}

}

//Function used to return a random response
func ElizaResponse(input string) string {

	//Set of responses
	response := [3]string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?"}
	//Random response index
	rIndex := rand.Intn(4-1) + 1
	//Return the response
	return response[rIndex-1]
}
