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

type SurfData struct {
	Year, Month, Day, Hour, Min int
	WVHT, SwH, SwP, WWH, WWP    float64
	SwD, WWD, STEEPNESS         string
	APD, MWD                    float64
	Count                       int
}

func main() {
	http.HandleFunc("/", surfDataRequest)
	http.ListenAndServe(":3000", nil)
}

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
