package main

import (
	"io/ioutil"
	"log"
	"net/http"
 )

 func main() {
	resp, err := http.Get("https://raw.githubusercontent.com/robconery/json-sales-data/master/data/customers.json")
	if err != nil {
	   log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	   log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
 }