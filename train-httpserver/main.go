package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Welcome to Go Rookie Gym :)")
	})
	http.HandleFunc("/json", handler)

	// ここの下に追加
}

type Payload struct {
	// jsonタグを追加してみよう
	// key名はname
	MyName string
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var p Payload
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	name := p.MyName
	p.MyName = fmt.Sprintf("Hello %s", name)
	j, err := json.Marshal(&p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
