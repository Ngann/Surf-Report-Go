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
	"math"
)

type SurfData struct {
	Year, Month, Day, Hour, Min int
	WVHT, SwH, SwP, WWH, WWP    float64
	SwD, WWD, STEEPNESS         string
	APD, MWD                    float64
	Count                       int
}

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

type Triangle struct {
	Base float64
	Height float64
}


type Shape interface {
	Area() float64
}

// Go interface resolution is implicit , If the type you pass in matches what the interface is asking for, it will compile.
// Rectangle has a method called Area that returns a float64 so it satisfies the Shape interface
// Circle has a method called Area that returns a float64 so it satisfies the Shape interface

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) *0.5
}

//type Shape interface has Area() method which then allows each shape to call it's own method.