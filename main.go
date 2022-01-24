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

	port := os.Getenv("PORT")

	http.HandleFunc("/", getAllTimezone)
	fmt.Println("Server started at port 8080")
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("Not Running Server...", err.Error())
	}
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
