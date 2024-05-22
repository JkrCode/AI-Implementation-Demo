package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func queryOllama(question string) (llamaResponse, error) {
	//prepare response
	response := llamaResponse{pid: -1, humanAnswer: "", tags: nil}
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
        return response, err
    }

    // Execute the request to the Ollama chat endpoint
    AiResponse, err := http.Post("http://localhost:11434/api/chat", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return response, err
    }
    defer AiResponse.Body.Close()

    // Read the response body
    body, err := io.ReadAll(AiResponse.Body)
    if err != nil {
        return response, err
    }

    // Parse the JSON response
    var llamaJSONresponse map[string]interface{}
    if err := json.Unmarshal(body, &llamaJSONresponse); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
        return response, err
    }

	fmt.Println(llamaJSONresponse)

    // Accessing the nested structure
    message, ok := llamaJSONresponse["message"].(map[string]interface{})
    if !ok {
        fmt.Println("Invalid response from Ollama: message field")
        return response , err
    }

    contentStr, ok := message["content"].(string)
    if !ok {
        fmt.Println("Invalid response from Ollama: content field")
        return response, err
    }

    // Unmarshal the JSON string in the content field
    var content map[string]interface{}
    if err := json.Unmarshal([]byte(contentStr), &content); err != nil {
        fmt.Println("Error unmarshaling content JSON:", err)
        return response , err
    }

    humanAnswer, ok := content["humanAnswer"].(string)
    if !ok {
        fmt.Println("Invalid response from Ollama: humanAnswer field")
        return response, err
    }

    pid, ok := content["pid"].(float64) // JSON numbers are decoded as float64 in Go
    if !ok {
        fmt.Println("Invalid response from Ollama: pid field")
        return response, err
    }

	tags, ok := content["tags"].(string) // JSON numbers are decoded as float64 in Go
    if !ok {
        fmt.Println("Invalid response from Ollama: pid field")
        return response, err
    }
	tagList := strings.Split(tags, " ")

	response.humanAnswer = humanAnswer
	response.pid = int(pid)
	response.tags = tagList

    return response, nil
}