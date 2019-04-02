package main

import (
	// "encoding/json"
	// "io/ioutil"
	"log"
	// "net/http"
	"regexp"
	// "strconv"
	"strings"
	// "fmt"
	// "reflect"
)

type SurfData struct {
	Year, Month, Day, Hour, Min int
	WVHT, SwH, SwP, WWH, WWP    float64
	SwD, WWD, STEEPNESS         string
	APD, MWD                    float64
	Count                       int
}

// var testRow = `
// 	2019 04 02 17 00  1.7  1.7 16.0  0.2  3.4 WNW  SW      SWELL  9.6 292
// `

// var testString = `
// 	#YY  MM DD hh mm WVHT  SwH  SwP  WWH  WWP SwD WWD  STEEPNESS  APD MWD
// 	#yr  mo dy hr mn    m    m  sec    m  sec  -  degT     -      sec degT
// 	2019 04 02 17 00  1.7  1.7 16.0  0.2  3.4 WNW  SW      SWELL  9.6 292
// 	2019 04 02 16 00  1.8  1.8 14.8  0.2  3.4 WNW NNW      SWELL  9.3 286
// 	2019 04 02 15 00  1.7  1.7 16.0  0.2  3.4 WNW NNE      SWELL  9.4 288
// 	2019 04 02 14 00  1.5  1.5 13.8  0.2  3.6 WNW NNW      SWELL  9.1 285
// 	2019 04 02 13 00  1.7  1.7 17.4  0.2  3.8 WNW NNE      SWELL 10.1 288
// 	2019 04 02 12 00  1.9  1.8 16.0  0.2  3.7 WNW NNW      SWELL 10.5 289
// `


func main() {
	// http.HandleFunc("/", surfDataRequest)
	// http.ListenAndServe(":3000", nil)
}

//the httprequest will return a response that is a string
// fmt.Println(reflect.TypeOf(surf)) is a string
// fmt.Println(reflect.TypeOf(rows)) is now an array of string because we split the data.

func ParseString(data string) {
	surf := string(data)
	//to match all leading/trailing whitespac
	leadSpace := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	//to match 2 or more whitespace symbols inside a string
	extraSpace := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	final := leadSpace.ReplaceAllString(surf, "")
	final = extraSpace.ReplaceAllString(final, " ")
	rows := strings.Split(final, "\n")
		log.Println(rows)
	
}

type Rectangle struct {
	Width float64
	Height float64
}


type Circle struct {
	Radius float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}