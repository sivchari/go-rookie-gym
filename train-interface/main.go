package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type user struct {
	id   int    `json:"id"`
	name string `json:"name"`
}

func (u *user) UnmarshalJSON(b []byte) error {
	u2 := &struct {
		ID   int
		Name string
	}{}
	if err := json.Unmarshal(b, &u2); err != nil {
		return err
	}
	u.id = u2.ID
	u.name = u2.Name
	return nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// defer r.Body.Close() http.Requestに関してはCloseがいらない
		var u user
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := json.Unmarshal(body, &u); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Println(u)
		w.WriteHeader(http.StatusOK)
	})
	log.Println(http.ListenAndServe(":8080", nil))
}
