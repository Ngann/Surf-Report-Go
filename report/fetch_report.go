package report

import (
	"io/ioutil"
	"log"
	"net/http"
)

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
