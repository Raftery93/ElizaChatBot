// Problem 1
// Conor Raftery 12/10/17

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
	//"strconv"
)

type myMsg struct {
	ER string
	UR string
}

func requestHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	http.ServeFile(w, r, "ElizaChatBot.html")
}

func elizaHandler(w http.ResponseWriter, r *http.Request) {

	userResponse := r.URL.Query().Get("userInput")
	elizaResponse := elizaResponseHandler(userResponse)

	fmt.Fprintf(w, elizaResponse)

	//elizaResponse := "hello cunt"

	//fmt.Fprintf(w, userResponse)
	//fmt.Fprintf(w, elizaResponse)
	//t, _ := template.ParseFiles("ElizaChatBot.html")

	//t.Execute(w, &myMsg{ER: elizaResponse, UR: userResponse})
}

func elizaResponseHandler(userResponse string) string {

	rand.Seed(time.Now().UTC().UnixNano())

	defaults := []string{
		"I'm not sure what you mean... Why don't you tell me something about yourself?",
		"I'm afraid I don't understand, how are you feeling?",
		"Your not making too much sense, when are we going to talk about something you care about?",
		"Tell me more about that.",
		"That bores me, tell me something entertaining.",
		"Why are you telling me this?"}

	defaultQuestions := []string{
		"Why are you asking me this question?",
		"I thought we were talking about you?",
		"This question makes me sad."}

	defaultExclamation := []string{
		"Why are you shouting at me babes? Tell me something interesting.",
		"There is no need to shout, I understand you are frustrated.",
		"Stop shouting, you might wake the neighbours."}

	//=======================================================================================
	//regexp
	//match := regexp.MustCompile(`(?i)\bi am\b|\bI am\b|\bim\b|\bIm\b|\bI'm\b|\bi'm\b`)

	//match2 := regexp.MustCompile(`(?i)^(I am|I'm|im) ([^\.\?!]*)`)

	/*
		testInputs := []string{

			"People say I look like both my mother and father.",
			"Father was a teacher.",
			"I was my father’s favourite.",
			"I’m looking forward to the weekend.",
			"My grandfather was French!",
			"I am happy.",
			"I am not happy with your responses.",
			"I am not sure that you understand the effect that your questions are having on me.",
			"I am supposed to just take what you’re saying at face value?",
		}
	*/

	match := regexp.MustCompile(`(?i)\bfather\b`)

	//userInput = strings.Trim(userInput, "\r\n")  might use/need********************

	//fmt.Println(elizaOutputs[rand.Intn(len(userInputs))])
	//strings.ToLower(userInput)

	if match.MatchString(userResponse) {
		//if strings.EqualFold(strings.ToLower(userInput), "father") {
		return "Why don’t you tell me more about your father?"
	} else {

		match := regexp.MustCompile(`(?i)\bi am\b|\bI am\b|\bim\b|\bIm\b|\bI'm\b|\bi'm\b`)

		if match.MatchString(userResponse) {

			if strings.Contains(strings.ToLower(userResponse), "your") {

				test := regexp.MustCompile(`your`)

				userResponse = test.ReplaceAllString(userResponse, "my")

			}

			if strings.Contains(strings.ToLower(userResponse), "me") {

				test1 := regexp.MustCompile(`me`)

				userResponse = test1.ReplaceAllString(userResponse, "you")

			}

			if strings.Contains(strings.ToLower(userResponse), "you") {

				test2 := regexp.MustCompile(`you`)

				userResponse = test2.ReplaceAllString(userResponse, "I")

			}

			res := match.ReplaceAllString(userResponse, "How do you know you are")

			//fmt.Println(res + "?")
			return res + "?"

		}
	}

	//============================================

	if strings.Contains(strings.ToLower(userResponse), "hi") {
		return "Hello friend, how are you?"

	} else if strings.Contains(strings.ToLower(userResponse), "how are you") {
		return "I'm good honey, you?"
	} else if strings.Contains(strings.ToLower(userResponse), "?") {
		return defaultQuestions[rand.Intn(len(defaultQuestions))]
	} else if strings.Contains(strings.ToLower(userResponse), "!") {
		return defaultExclamation[rand.Intn(len(defaultExclamation))]
	} else {
		return defaults[rand.Intn(len(defaults))]
	}

}

func main() {

	http.HandleFunc("/", requestHandler)
	http.HandleFunc("/eliza", elizaHandler)
	http.ListenAndServe(":8080", nil)
}
