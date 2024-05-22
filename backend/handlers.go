package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getAiResponseHandler(w http.ResponseWriter, r *http.Request) {
    promptBuffer, err := io.ReadAll(r.Body)
    if err != nil {
        return
    }
    defer r.Body.Close()

    var prompt map[string]interface{}
    if err := json.Unmarshal(promptBuffer, &prompt); err != nil {
        return
    }

    question := prompt["content"].(string)

	fmt.Println("endpoint hit")
	
    answer,pid ,err := queryOllama(question)
    if err != nil {
        fmt.Fprintln(w, "Error querying Ollama:", err)
        return
    }
	pString := ""
	if pid==2{
		response, err := http.Get("http://localhost:3033")
		if err != nil {
			return 
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return 
		}
		responseString := string(body)

		pString = pString + responseString
	}
    fmt.Fprintln(w, answer + " " + pString)
}