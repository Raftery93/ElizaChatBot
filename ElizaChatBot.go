// Problem 1
// Conor Raftery 12/10/17

package main

import (
	"html/template"
	"net/http"
	//"strconv"

)

type myMsg struct {
	ER string
	UR string
	//GuessMessage string
	//Guess int
}


func requestHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type","text/html")

	http.ServeFile(w, r, "ChatHomePage.html")
}

func elizaHandler(w http.ResponseWriter, r *http.Request){


	elizaResponse :="Hello User, how are you?"

	userResponse := r.FormValue("userResponse")


	


	t, _ := template.ParseFiles("ElizaChatBot.html")

	t.Execute(w, &myMsg{ER:elizaResponse, UR:userResponse})
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.HandleFunc("/guess", elizaHandler)
	http.ListenAndServe(":8080", nil)
}