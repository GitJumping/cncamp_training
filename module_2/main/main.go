package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/profile", func(writer http.ResponseWriter, request *http.Request) {
		data := map[string]string{
			"username": "小明",
			"email":    "xiaoming@163.com",
		}
		err := JSON(writer, data)
		check(err)
	})
	http.ListenAndServe(":8080", nil)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func JSON(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	return encoder.Encode(data)
}
