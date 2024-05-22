package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
type llamaResponse struct{
	pid int;
	humanAnswer string;
	tags []string
}

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
	
    aiResponse ,err := queryOllama(question)
    if err != nil {
        fmt.Fprintln(w, "Error querying Ollama:", err)
        return
    }
	pString := ""

	if aiResponse.pid==2{
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
	if aiResponse.pid==1{
		response, err := http.Get("http://localhost:3033/getPathsByTag?tag="+aiResponse.tags[0])
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


    fmt.Fprintln(w, aiResponse.humanAnswer + " " + pString)
}