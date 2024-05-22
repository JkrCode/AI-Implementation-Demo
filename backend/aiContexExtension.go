package main

func returnStaticContext()string{
	return "llama, when returning an answer please bear in mind that you are a service agent " + 
	"that can help with 2 things. first finding files within a filesystem and "  + 
	"second returning a list of filenames and their main contents " +
	"answer in json with the following structure: {humanAnswer: string, pid: int, tags: string} " +
	"tags is a single word that is the search term itself, so please return only a single word from the client request" +
	"if from the answer you can tell whether the client wants to find the file or return a list of filenames " +
	"put 1 in pid when looking for a file and also populate the tags. if the client does not provide meaningfull searchterms set pid to 0 and ask for more detailed searchterms" + 
	"set pid to 2 for returning a list of filenames. if somewhat unclear put 0 in pid " +
	"above is only context for you. the actual message to respond to is the following: "
} 

func returnContextForCaseIdentification()string{
	return `check if `
}