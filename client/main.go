package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Second}
	req, err := http.NewRequest(`POST`, `http://localhost:5000/Olena`, nil)
	if err != nil {
		fmt.Printf("Error: %s\\n", err)
		return
	}
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("Error: %s\\n", err)
		return
	}

	// read response body
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		fmt.Println(error)
	}
	// print response body
	log.Println(string(body))

	defer resp.Body.Close()
}
