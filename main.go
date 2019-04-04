package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "https://www.ndbc.noaa.gov/data/realtime2/46029.spec"
	fetch := SurfDataRequest(url)
	fmt.Println(fetch)

}

func SurfDataRequest(url string) string {
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("StatusCode:", res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}
