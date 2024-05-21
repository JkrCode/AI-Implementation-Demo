import React, { useState } from "react";
import "./App.css";

function App() {
  const [inputValue, setInputValue] = useState("");
  const [response, setResponse] = useState("");

  const handleInputChange = (event) => {
    setInputValue(event.target.value);
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    console.log("Form submitted");

    try {
      console.log("Sending request to server");
      const res = await fetch("http://localhost:8081/api", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ content: inputValue }), // Send JSON data
      });

      console.log("Request sent, awaiting response");
      if (!res.ok) {
        throw new Error(`HTTP error! status: ${res.status}`);
      }

      console.log("Response received");
      const data = await res.text();
      console.log("Response data:", data);
      setResponse(data);
    } catch (error) {
      console.error("Error fetching data:", error);
      setResponse("Error fetching data");
    }
  };

  return (
    <div>
      <p>Put in some text here to chat with our service bot</p>
      <form onSubmit={handleSubmit}>
        <input value={inputValue} onChange={handleInputChange}></input>
        <button type="submit">Send text</button>
      </form>
      <div>
        <p>Response:</p>
        <p>{response}</p>
      </div>
    </div>
  );
}

export default App;
