package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type RandomDogResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>DANTE</h1>")
}

func getRandomDog() RandomDogResponse {
	url := "https://dog.ceo/api/breeds/image/random"
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
		return RandomDogResponse{
			Status:  "Failed",
			Message: "Random dog serer not available",
		}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var parsedBody RandomDogResponse
	json.Unmarshal(body, &parsedBody)

	return parsedBody
}

func randomDogHandler(w http.ResponseWriter, r *http.Request) {
	randomDogResponse := getRandomDog()
	if randomDogResponse.Status != "Failed" {
		fmt.Fprintf(w, "<div style=display:flex;justify-self:center;width:500px;height:500px;><img style=height:-webkit-fill-available;width:-webkit-fill-available; src=%s /> </div>", randomDogResponse.Message)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/doggo", randomDogHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
