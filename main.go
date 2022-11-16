package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getAllTimezone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	data := loadData()
	w.Write(data)
}

func main() {
	http.HandleFunc("/", getAllTimezone)
	fmt.Println("running server")
	log.Fatal(http.ListenAndServe(":7001", nil))
}

func loadData() []byte {
	jsonFile, err := os.Open("./response.json")

	if err != nil {
		log.Println("Not found File", err.Error())
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Println("Not Read File", err.Error())
	}

	return data

}
