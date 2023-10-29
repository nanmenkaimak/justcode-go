package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const portNumber = ":8081"

func main() {
	http.HandleFunc("/get_from_first", getDataFromFirstServer)
	http.HandleFunc("/message", message)
	log.Fatal(http.ListenAndServe(portNumber, nil))
}

func getDataFromFirstServer(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8080/message")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(body)
}

func message(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from second server"))
}
