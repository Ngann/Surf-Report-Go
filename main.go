package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var sampleRow = `
	2019 04 02 17 00  1.7  1.7 16.0  0.2  3.4 WNW  SW      SWELL  9.6 292
`

// #YY  MM DD hh mm WVHT  SwH  SwP  WWH  WWP SwD WWD  STEEPNESS  APD MWD

type SurfData struct {
	Year, Month, Day, Hour, Min int
	WVHT, SwH, SwP, WWH, WWP    float64
	SwD, WWD, STEEPNESS         string
	APD, MWD                    float64
}

func main() {
	// k := removeSpacesFromString(sampleRow)
	m := removeSpaces(sampleRow)
	// m := parseStringCount(k)
	// n := parseString(k)
	o := getRowDataWWD(m)
	// fmt.Println(sampleString)
	// fmt.Println(k)
	fmt.Println(m)
	fmt.Println(o)
}

func removeSpaces(data string) []string {
	surf := data
	newString := strings.Fields(surf)
	return newString
}

func removeSpacesFromString(data string) string {
	surf := data
	leadSpace := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	extraSpace := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	newString := leadSpace.ReplaceAllString(surf, "")
	newString = extraSpace.ReplaceAllString(newString, " ")

	return newString
}

func parseStringCount(data string) int {
	row := strings.Split(data, " ")
	return len(row)
}

func parseString(data string) []string {
	row := strings.Split(data, " ")
	return row
}

func getRowDataWWD(data []string) string {
	WWD := string(data[11])
	return WWD
}

func getRowDataSwH(data []string) float64 {
	SwH, _ := strconv.ParseFloat(data[6], 64)
	return SwH
}
func getRowDataSwP(data []string) float64 {
	SwP, _ := strconv.ParseFloat(data[7], 64)
	return SwP
}

func windScore(data string) float64 {
	if data == "E" {
		return 5
	} else if data == "NE" {
		return 4
	} else if data == "SE" {
		return 4
	} else if data == "S" {
		return 3
	} else {
		return 1
	}
}

func swellScore(data float64) float64 {
	if data >= 16 {
		return 5
	} else if data >= 12 {
		return 4
	} else if data >= 10 {
		return 3
	} else {
		return 1
	}
}

func waveScore(swp, swh float64) float64 {
	wave := swp * swh
	if wave > 30 {
		return 5
	} else if wave >= 25 {
		return 4
	} else if wave >= 20 {
		return 3
	} else if wave >= 11 {
		return 2
	} else {
		return 1
	}
}

// func averageScore(wind, swell, wave float64) float64 {
// 	average := (wind + swell + wave) / 3
// 	return average
// }

func getScoreForToday(data []string) float64 {
	return 1
}
