package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

//Structure for a category
type Category struct {
	pattern  string   //Pattern to match
	response []string //Response to chose from
}

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

	//Lop the categories
	for i, req := range request {
		fmt.Println(i+1, "\nRequest:", req)
		//Get a response
		if response := getResponse(req); response != "" {
			fmt.Println("Response:", response, "\n")
		}
	}

}

func getResponse(request string) string {
	//Steh of patterns and maching answers
	categories := [2]Category{
		Category{"father(.*)", []string{"Why don’t you tell me more about your father?"}},
		Category{"(.*)", []string{
			"I’m not sure what you’re trying to say. Could you explain it to me?",
			"How does that make you feel?",
			"Why do you say that?"}},
	}

	//Lop the categories
	for _, cat := range categories {
		//Try to get an answer for this category
		if found, answer := ElizaResponse(cat, request); found == true {
			//Return the answer
			return answer
		}
	}
	//Return an empty string
	return ""
}

//Function used to return a random response
func ElizaResponse(category Category, input string) (bool, string) {
	//Compile the pattern
	var pattern = regexp.MustCompile("(?i)" + category.pattern)
	//If there was a match
	if matched := pattern.MatchString(input); matched == true {
		//Index of the answer
		rIndex := 1
		//Check if there is more than one answer
		if len(category.response) > 1 {
			//Random response index
			rIndex = rand.Intn(len(category.response)-1) + 1
		}
		//Return the response
		return true, category.response[rIndex-1]
	}

	//Return false
	return false, ""

}
