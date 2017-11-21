// Problem 1
// Conor Raftery 12/10/17

package main

import (
	"fmt"
	"net/http"
	"strings"
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

	if strings.Contains(strings.ToLower(userResponse), "hi") {
		println("in first if")
		elizaResponse := "if 1"
		return elizaResponse

	} else if strings.Contains(strings.ToLower(userResponse), "no") {
		println("in second if")
		elizaResponse := "if 2"
		return elizaResponse
	} else {
		println("in else")
		elizaResponse := "if 3"
		return elizaResponse
	}

}

func main() {

	http.HandleFunc("/", requestHandler)
	http.HandleFunc("/eliza", elizaHandler)
	http.ListenAndServe(":8080", nil)
}
