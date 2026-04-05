package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	if len(params.Body) <= 140 {
		type validChirp struct {
			Valid bool `json:"valid"`
		}
		respondWithJSON(w, 200, validChirp{Valid: true})
		return
	}
	respondWithErr(w, 400, "Chirp is too long.")
}