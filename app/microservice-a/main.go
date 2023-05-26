// main.go (Микросервис A)
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Выполнение запроса к Микросервису B
	response, err := http.Get("http://microservice-b:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Чтение ответа от Микросервиса B
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Отправка ответа клиенту
	fmt.Fprintf(w, "Response from Microservice A\n")
	fmt.Fprintf(w, "Response from Microservice B: %s\n", body)
}
