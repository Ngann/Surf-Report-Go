package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var sampleRow = `
	2019 04 02 17 00  1.7  1.7 16.0  0.2  3.4 WNW  SW      SWELL  9.6 292
`

var sampleString = `#YY  MM DD hh mm WVHT  SwH  SwP  WWH  WWP SwD WWD  STEEPNESS  APD MWD
#yr  mo dy hr mn    m    m  sec    m  sec  -  degT     -      sec degT
2019 04 04 16 00  1.1  0.8 14.8  0.6  5.0 WSW WSW      SWELL  5.3 241
2019 04 04 15 00  1.0  0.8 14.8  0.6  5.3 WSW WSW      SWELL  5.4 243
2019 04 04 14 00  1.1  1.0 14.8  0.5  4.2   E WSW      SWELL  6.1  80
2019 04 04 13 00  1.2  1.1 14.8  0.4  3.8 SSW WSW      SWELL  5.8 211
2019 04 04 12 00  1.3  1.2 14.8  0.4  3.8  SW WSW      SWELL  5.8 218
2019 04 04 11 00  1.2  1.1  6.2  0.4  4.0 WNW WSW      STEEP  5.6 282
2019 04 04 10 00  1.3  1.2 13.8  0.4  3.4  SW  SW      SWELL  6.0 236
2019 04 04 09 00  1.3  1.3  7.1  0.4  4.0 WSW WSW      STEEP  5.9 250
2019 04 04 08 00  1.4  1.3  5.6  0.4  3.8 WSW   W VERY_STEEP  5.9 256
2019 04 04 07 00  1.4  1.4  6.2  0.4  3.4   W  SW      STEEP  5.8 267`

type SurfData struct {
	Year, Month, Day, Hour, Min int
	WVHT, SwH, SwP, WWH, WWP    float64
	SwD, WWD, STEEPNESS         string
	APD, MWD                    float64
}

func main() {
	http.HandleFunc("/", surfDataRequest)
	http.ListenAndServe(":3000", nil)
	// a := parseString(sampleString)
	// b := getSurfData(a)
	// fmt.Println(b)
	// fmt.Println(len(b))
}

func parseString(buoyData string) []string {
	surf := buoyData
	leadSpace := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	extraSpace := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	final := leadSpace.ReplaceAllString(surf, "")
	final = extraSpace.ReplaceAllString(final, " ")
	rows := strings.Split(final, "\n")

	return rows
}

func createSurfStruct(row []string) SurfData {
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

	return SurfData{
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
	}
}

func getSurfData(rows []string) []SurfData {
	var allSurfData []SurfData
	fmt.Println(len(rows))
	for i := 2; i < len(rows); i++ {
		row := strings.Split(rows[i], " ")
		if len(row) < 2 {
			continue
		}
		allSurfData = append(allSurfData, createSurfStruct(row))
	}
	return allSurfData
}

func surfDataRequest(w http.ResponseWriter, r *http.Request) {
	// Set the content-type header so clients know to expect json
	// Originally when running application on the front end, ran into some CORS issue related to "No 'Access-Control-Allow-Origin", so we set the header to allow access of the data from 3000 to be render on 3001
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

	parseReposnse := parseString(string(body))
	buildStruct := getSurfData(parseReposnse)

	// create json data
	js, err := json.Marshal(buildStruct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}
