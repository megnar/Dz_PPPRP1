// main.go (Микросервис B)
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
	// Выполнение запроса к публичному API
	response, err := http.Get("https://yandex.ru/pogoda/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Чтение ответа от публичного API
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Отправка ответа клиенту
	fmt.Fprintf(w, "Response from Microservice B\n")
	fmt.Fprintf(w, "Response from External API: %s\n", body)
}
