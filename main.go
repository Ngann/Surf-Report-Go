package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

// structs are collection of fields
type SurfData struct {
	Year, Month, Day, Hour, Min int
	WVHT, SwH, SwP, WWH, WWP    float64
	SwD, WWD, STEEPNESS         string
	APD, MWD                    float64
	Count                       int
}

// The main function begins with a call to the http.HandleFunc, which tells the http package to handle all request to the web root("/") with function surfDataRequest
// It then calls the http.ListenAndServe, which specify that is should listen on port 3000.
func main() {
	http.HandleFunc("/", surfDataRequest)
	http.ListenAndServe(":3000", nil)
}

// The function surfDataRequest is a type of http.HandleFunc, and it takes an http.ResponseWritter and an http.Request as it's arguments.
// An http.ResponseWriter assembles the HTTP server's response; by writting to it, we send the data to the HTTP client. (What is the http client?)
// An http.Request is a data structure that represents the client HTTP server's reponse.

//example, below will render the request from function viewHandler on localhost:8080/view/ etc...
// func main() {
//     http.HandleFunc("/view/", viewHandler)
//     http.HandleFunc("/edit/", editHandler)
//     http.HandleFunc("/save/", saveHandler)
//     log.Fatal(http.ListenAndServe(":8080", nil))
// }

func surfDataRequest(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3001")
	if (*r).Method == "OPTIONS" {
		return
	}
	var err error

	// request http api
	res, err := http.Get("https://www.ndbc.noaa.gov/data/realtime2/46029.spec")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("StatusCode:", res.StatusCode)

	// read body
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	surf := string(body)
	//to match all leading/trailing whitespac
	leadSpace := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	//to match 2 or more whitespace symbols inside a string
	extraSpace := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	final := leadSpace.ReplaceAllString(surf, "")
	final = extraSpace.ReplaceAllString(final, " ")
	rows := strings.Split(final, "\n")

	var allSurfData []SurfData

	for i := 2; i < 100; i++ {
		row := strings.Split(rows[i], " ")
		if len(row) < 2 {
			continue
		}
		year, _ := strconv.Atoi(row[0])
		month, _ := strconv.Atoi(row[1])
		day, _ := strconv.Atoi(row[2])
		hour, _ := strconv.Atoi(row[3])
		min, _ := strconv.Atoi(row[4])
		WVHT, _ := strconv.ParseFloat(row[5], 64)
		SwH, _ := strconv.ParseFloat(row[6], 64)
		SwP, _ := strconv.ParseFloat(row[7], 64)
		WWH, _ := strconv.ParseFloat(row[8], 64)
		WWP, _ := strconv.ParseFloat(row[9], 64)
		SwD := row[10]
		WWD := row[11]
		STEEPNESS := row[12]
		APD, _ := strconv.ParseFloat(row[13], 64)
		MWD, _ := strconv.ParseFloat(row[14], 64)
		Count := (i - 1)
		var surfData = SurfData{
			year,
			month,
			day,
			hour,
			min,
			WVHT,
			SwH,
			SwP,
			WWH,
			WWP,
			SwD,
			WWD,
			STEEPNESS,
			APD,
			MWD,
			Count,
		}

		allSurfData = append(allSurfData, surfData)
	}
	// log.Println(allSurfData[0])

	// create json data
	js, err := json.Marshal(allSurfData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}
