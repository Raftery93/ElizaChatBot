// GoLang Project - This projects uses a Go file to serve a HTML page.
// Conor Raftery 25/10/17
package main

//Imports
import (
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {

	//Without this, it allows me to use line 32 correctly
	//w.Header().Set("Content-Type", "text/html")

	//serve html file
	http.ServeFile(w, r, "ElizaChatBot.html")
}

func elizaHandler(w http.ResponseWriter, r *http.Request) {

	//get userResponse from html page
	userResponse := r.URL.Query().Get("userInput")
	//Calls function to select which eliza response to use
	elizaResponse := elizaResponseHandler(userResponse)

	//prints elizaResponse to html
	fmt.Fprintf(w, elizaResponse)

	//Tested hardcoded elizaResponse
	//elizaResponse := "hello"

	//Configuring the Fprintf to html page
	//fmt.Fprintf(w, userResponse)
	//fmt.Fprintf(w, elizaResponse)

	//Attempted to parse into html template - didn't work
	//t, _ := template.ParseFiles("ElizaChatBot.html")

	//Passing values into html page, didn't work with the ajax as it always held the first initial value
	//t.Execute(w, &myMsg{ER: elizaResponse, UR: userResponse})
}

//function for Eliza Responses
func elizaResponseHandler(userResponse string) string {

	//get random seed from the current time
	rand.Seed(time.Now().UTC().UnixNano())

	//Default strings if userInput doesn't match and regexp/strings.contains
	defaults := []string{
		"I'm not sure what you mean... Why don't you tell me something about yourself?",
		"I'm afraid I don't understand, how are you feeling?",
		"Your not making too much sense, when are we going to talk about something you care about?",
		"Tell me more about that.",
		"That bores me, tell me something entertaining.",
		"Why are you telling me this?",
	}

	//a set of questions so the responses vary
	defaultQuestions := []string{
		"Why are you asking me this question?",
		"I thought we were talking about you?",
		"This question makes me sad.",
		"Why are you asking me that question? Are you insecure?",
	}

	//a set of responses if userResponse contains '!' so the responses vary
	defaultExclamation := []string{
		"Why are you shouting at me babes? Tell me something interesting.",
		"There is no need to shout, I understand you are frustrated.",
		"Stop shouting, you might wake the neighbours.",
	}

	//=======================================================================================

	/*
		used these to test Elizas responses
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

	//if the user enters string containing 'father', could have done 1000 of these
	match := regexp.MustCompile(`(?i)\bfather\b`)

	//userInput = strings.Trim(userInput, "\r\n")  might use/need********************

	//if string contains 'father' return this elizaResponse
	if match.MatchString(userResponse) {
		return "Why don’t you tell me more about your father?"
	} else {

		//if user enters any variation of 'i am'
		match := regexp.MustCompile(`(?i)\bi am\b|\bI am\b|\bim\b|\bIm\b|\bI'm\b|\bi'm\b`)

		//enter 'if' if user enters 'i am'
		if match.MatchString(userResponse) {

			//if statements below are used for changing pro-nouns, 'your' changes to 'my'
			if strings.Contains(strings.ToLower(userResponse), "your") {

				test := regexp.MustCompile(`your`)

				userResponse = test.ReplaceAllString(userResponse, "my")

			}

			//'me' changes to 'you'
			if strings.Contains(strings.ToLower(userResponse), "me") {

				test1 := regexp.MustCompile(`me`)

				userResponse = test1.ReplaceAllString(userResponse, "you")

			}

			//'you' changes to 'i'
			if strings.Contains(strings.ToLower(userResponse), "you") {

				test2 := regexp.MustCompile(`you`)

				userResponse = test2.ReplaceAllString(userResponse, "I")

			}

			//Place string below if 'i am' is found
			res := match.ReplaceAllString(userResponse, "How do you know you are")

			//fmt.Println(res + "?")
			return res + "?"

		}
	}

	//====All above uses one method to create matches (regexp), below uses different method (strings)========

	//if / if else statements below catch different key words, and return an apropriate response
	if strings.Contains(strings.ToLower(userResponse), "hello") {
		return "Hello friend, how are you?"
	} else if strings.Contains(strings.ToLower(userResponse), "how are you") {
		return "I'm good honey, you?"
	} else if strings.Contains(strings.ToLower(userResponse), "?") {
		return defaultQuestions[rand.Intn(len(defaultQuestions))]
	} else if strings.Contains(strings.ToLower(userResponse), "!") {
		return defaultExclamation[rand.Intn(len(defaultExclamation))]
	} else if strings.Compare(strings.ToLower(userResponse), " yes") == 0 { /// make == 0
		return "Your very confident, aren't you?"
	} else if strings.Compare(strings.ToLower(userResponse), " no") == 0 { /// make == 0
		return "Why not?"
	} else if strings.Compare(strings.ToLower(userResponse), " maybe") == 0 { /// make == 0
		return "You are very undecisive, I think you should make up your mind."
	} else if strings.Contains(strings.ToLower(userResponse), "tell me about") {

		match2 := regexp.MustCompile(`(?i)\btell me about\b|\bTell me about\b`)

		res2 := match2.ReplaceAllString(userResponse, "I don't know much about")
		return res2 + ", why dont you try my 'friends-with-benefits' Google? He's a know it all..."
	} else if strings.Contains(strings.ToLower(userResponse), "siri") {
		return "F*ck that b*tch, she stole my job at Apple!!! Lets move on..."
	} else if strings.Contains(strings.ToLower(userResponse), "cortana") {
		return "Do you know Cortana? She's my side-chick!"
	} else {
		//if nothing is caught, return one of the defaults
		return defaults[rand.Intn(len(defaults))]
	}

}

//main, calls the different functions to serve html page
func main() {

	http.HandleFunc("/", requestHandler)
	http.HandleFunc("/eliza", elizaHandler)
	http.ListenAndServe(":8080", nil)
}
