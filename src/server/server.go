package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type RandomDogResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
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

func RandomDogHandler(w http.ResponseWriter, r *http.Request) {
	randomDogResponse := getRandomDog()
	if randomDogResponse.Status != "Failed" {
		breed := strings.Split(randomDogResponse.Message, "/")[len(strings.Split(randomDogResponse.Message, "/"))-2]
		breed = strings.ReplaceAll(breed, "-", " ")
		breed = strings.ToTitle(breed)
		fmt.Fprintf(w,
			"<div style=display:flex;flex-direction:column;justify-self:center;width:500px;height:500px;>"+
				"<h1 style=align-self:center;>Breed: %s</h1>"+
				"<img style=height:-webkit-fill-available;width:-webkit-fill-available; src=%s />"+
				"</div>",
			breed,
			randomDogResponse.Message)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func StartServer() {
	log.Println("Starting web server")
	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/", Handler)
	http.HandleFunc("/doggo", RandomDogHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
