package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Chirp struct {
	Body string `json:"body"`
}


func respondWithErr(w http.ResponseWriter, code int, msg string) {
	type chirpErr struct {
		Error string `json:"error"`
	}
	
	data, err := json.Marshal(chirpErr{Error: msg})
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
	
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)

}

func (cfg *apiConfig) handleValidateChirp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := Chirp{}
	err := decoder.Decode(&params)
	if err != nil { 
		respondWithErr(w, 500, fmt.Sprintf("Error marshalling JSON: %s", err))
		return
	}
	body := params.Body
	if len(body) > 140 {
		type validChirp struct {
			Valid bool `json:"valid"`
		}
		respondWithErr(w, 400, "Chirp is too long.")	
		return
	}
	var cleanedBody []string
	for _, word := range strings.Split(body, " ") {
		switch strings.ToLower(word) {
		case "kerfuffle", "sharbert", "fornax":
			cleanedBody = append(cleanedBody, "****")
		default:
			cleanedBody = append(cleanedBody, word)
		}

	}
	respondWithJSON(w, 200, struct{Body string `json:"cleaned_body"`}{Body: strings.Join(cleanedBody, " ")})
}