package main

import "fmt"

func main() {

	pronobreSpAInges := map[string]string{
		"yo":       "i",
		"tu":       "you",
		"el":       "he",
		"ella":     "she",
		"eso":      "it",
		"ustedes":  "you",
		"nosotros": "we",
		"ellos":    "they",
	}

	delete(pronobreSpAInges, "yo")

	for ep, ingles := range pronobreSpAInges {
		fmt.Println(ep, " = ", ingles)
	}
}
