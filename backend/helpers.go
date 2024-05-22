package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func queryOllama(question string) (string, int, error) {
    // Create request payload with conversation context
    payload := map[string]interface{}{
        "model": "llama3",
        "messages": []map[string]string{
            {
                "role":    "user",
                "content":  returnStaticContext() + question,
            },
        },
		"format":"json",
        "stream": false,
    }
    jsonData, err := json.Marshal(payload)
    if err != nil {
        return "", -1, err
    }

    // Execute the request to the Ollama chat endpoint
    response, err := http.Post("http://localhost:11434/api/chat", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return "",-1, err
    }
    defer response.Body.Close()

    // Read the response body
    body, err := io.ReadAll(response.Body)
    if err != nil {
        return "",-1, err
    }

    // Parse the JSON response
    var llamaJSONresponse map[string]interface{}
    if err := json.Unmarshal(body, &llamaJSONresponse); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
        return "",-1, err
    }

	fmt.Println(llamaJSONresponse)

    // Accessing the nested structure
    message, ok := llamaJSONresponse["message"].(map[string]interface{})
    if !ok {
        fmt.Println("Invalid response from Ollama: message field")
        return "",-1 , err
    }

    contentStr, ok := message["content"].(string)
    if !ok {
        fmt.Println("Invalid response from Ollama: content field")
        return "" ,-1, err
    }

    // Unmarshal the JSON string in the content field
    var content map[string]interface{}
    if err := json.Unmarshal([]byte(contentStr), &content); err != nil {
        fmt.Println("Error unmarshaling content JSON:", err)
        return "" ,-1 , err
    }

    humanAnswer, ok := content["humanAnswer"].(string)
    if !ok {
        fmt.Println("Invalid response from Ollama: humanAnswer field")
        return "" ,-1, err
    }

    pid, ok := content["pid"].(float64) // JSON numbers are decoded as float64 in Go
    if !ok {
        fmt.Println("Invalid response from Ollama: pid field")
        return "" ,-1, err
    }

    return humanAnswer,int(pid), nil
}