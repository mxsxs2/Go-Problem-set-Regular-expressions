package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
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
	request := [9]string{
		"People say I look like both my mother and father.",
		"Father was a teacher.",
		"I was my father’s favourite.",
		"I’m looking forward to the weekend.",
		"My grandfather was French!",
		"I am happy.",
		"I am not happy with your responses.",
		"I am not sure that you understand the effect that your questions are having on me.",
		"I am supposed to just take what you’re saying at face value?"}

	//Lop the categories
	for _, req := range request {
		fmt.Println("\nRequest:", req)
		//Get a response
		if response := ElizaResponse(req); response != "" {
			fmt.Println("Response:", response)
		}
	}

}

//ElizaResponse is used to get a matching answer for an input string
func ElizaResponse(request string) string {
	//Steh of patterns and maching answers
	categories := [3]Category{
		Category{`.*\bfather\b.*`, []string{"Why don’t you tell me more about your father?"}},
		Category{`I am ([^.?!].*)[.?!]`, []string{"How do you know you are _?"}},
		Category{`(.*)`, []string{
			"I’m not sure what you’re trying to say. Could you explain it to me?",
			"How does that make you feel?",
			"Why do you say that?"}},
	}

	//Lop the categories
	for _, cat := range categories {
		//Try to get an answer for this category
		if found, answer := getResponse(cat, request); found == true {
			//Return the answer
			return answer
		}
	}
	//Return an empty string
	return ""
}

//getResponse is used to match a category to a given string and return a random answer from the categories list
func getResponse(category Category, input string) (bool, string) {
	//Compile the pattern
	var pattern = regexp.MustCompile(`(?i)` + category.pattern)
	//fmt.Println(category.pattern, input)
	//If there was a match
	if pattern.MatchString(input) {
		//Index of the answer
		rIndex := 1
		//Check if there is more than one answer
		if len(category.response) > 1 {
			//Random response index
			rIndex = rand.Intn(len(category.response)-1) + 1
		}

		//Check if response contains _
		if strings.Contains(category.response[rIndex-1], "_") {
			//Return the response
			return true, strings.Replace(category.response[rIndex-1], "_", strings.TrimSpace(pattern.FindStringSubmatch(input)[1]), 1)
		} else {
			//Return the response
			return true, category.response[rIndex-1]
		}
	}

	//Return false
	return false, ""
}
