package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func callMicro3(writer http.ResponseWriter, request *http.Request) {

	response, err := http.Get("http://micro3:8082/call")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
	cuerpo := "Micro2 que llama a " + string(body)
	_, err = writer.Write([]byte(cuerpo))
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	http.HandleFunc("/call", callMicro3)
	err := http.ListenAndServe(":8081", nil)
	log.Fatal(err)
}
