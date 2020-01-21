package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type LOTRCharacter struct {
	ID     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	People string `json:"people,omitempty"`
	Role   string `json:"role,omitempty"`
}

var characters []LOTRCharacter

func main() {
	http.HandleFunc("/api/lotr/characters/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		idRaw := parts[4]
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			log.Println("Could not convert ID to integer")
			return
		}
		for _, c := range characters {
			if id == c.ID {
				w.Header().Add("Content-Type", "application/json")
				data, err := json.Marshal(c)
				if err != nil {
					log.Println("Could not marshal character to JSON")
					return
				}
				w.Write(data)
				return
			}
		}
		http.NotFound(w, r)
	})

	f, err := os.Open("./LOTRcharacters.json")

	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	f.Close()

	err = json.Unmarshal(data, &characters)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8000", nil))
}
