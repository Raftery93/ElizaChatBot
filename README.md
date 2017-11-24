# Conor Raftery (G00274094) - Eliza Chatbot program
This project was part of my 3rd module, Data Representation and Querying. The project had to be written in GoLang, implementing AJAX. This in turn, serves a HTML file to a web browser. It uses javascript and json. The HTML page is accessed by running the Go WebApp in your comand prompt, and going to your local host in your web browser (127.0.0.1:8080). My projects contains two files, the GoLang file and the HTML file which also contains my JS/JQuery/JSON and AJAX.

#### What is an Eliza Chatbot?
ELIZA is an early natural language processing computer program created from 1964 to 1966[1] at the MIT Artificial Intelligence Laboratory by Joseph Weizenbaum.[2] Created to demonstrate the superficiality of communication between man and machine, Eliza simulated conversation by using a 'pattern matching' and substitution methodology that gave users an illusion of understanding on the part of the program, but had no built in framework for contextualizing events.[3] Directives on how to interact were provided by 'scripts', written originally in MAD-Slip, which allowed ELIZA to process user inputs and engage in discourse following the rules and directions of the script. The most famous script, DOCTOR, simulated a Rogerian psychotherapist and used rules, dictated in the script, to respond with non-directional questions to user inputs. As such, ELIZA was one of the first chatterbots, but was also regarded as one of the first programs capable of passing the Turing Test.
Taken from: [Wikipedia ELIZA](https://en.wikipedia.org/wiki/ELIZA)

#### Run the Program
1.Clone the repository: 
```
git clone https://github.com/Raftery93/ElizaChatBot.git
```
There is two ways to run this program
2. Navigate to the folder where you downloaded the file, and using your command prompt, enter the following: 
```
go build ElizaChatBot.go
```
This will create an executable file. To run this file, type into your command prompt:
```
ElizaChatBot.exe
```
3. Next, because this is a webApp, open your prefered web browser and navigate to the url(localhost):
```
127.0.0.1:8080
```

### Using the Program
This Eliza program will always respond with a certain response. If it recognises a word/statement, it will reply with a relative response. See what happens when you type (not case sensative):
Hello
* I am happy
* how are you
* 'yes', 'no' or 'maybe'
* mention the word 'Siri' or 'Cortana'
* Ask it any question followed by a '?' e.g. "What are you?"
* Try shouting at it followed by a '!' e.g "Shut up!"
* Ask it to tell you about something. e.g "Tell me about GMIT"
* It also incorporates pro-nouns, try using the following sentence for an example: "I am not sure that you understand the effect that your questions are having on me."

This is done by using imports 'regexp' and 'strings'


### Features
I attempted to incorporate a few features into the program, but I was unsuccesful. These features included a background image and audio on start up. (See code)
Sfter looking into the issue, it may be due to the fact that the browser does not support it.