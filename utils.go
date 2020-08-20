package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func nodeStatus() {
	method := "{\"method\":\"nodestatus\"}"
	var ip = "http://localhost:4009"

	bytesRepresentation, err := json.Marshal(method)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(ip, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}
	var result = ""

	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result)
}
