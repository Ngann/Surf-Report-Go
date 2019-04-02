package main

import (
	"testing"
	"regexp"
	"strings"
	"reflect"
)

var testRow = `
	2019 04 02 17 00  1.7  1.7 16.0  0.2  3.4 WNW  SW      SWELL  9.6 292
`

var testString = `
	#YY  MM DD hh mm WVHT  SwH  SwP  WWH  WWP SwD WWD  STEEPNESS  APD MWD
	#yr  mo dy hr mn    m    m  sec    m  sec  -  degT     -      sec degT
	2019 04 02 17 00  1.7  1.7 16.0  0.2  3.4 WNW  SW      SWELL  9.6 292
	2019 04 02 16 00  1.8  1.8 14.8  0.2  3.4 WNW NNW      SWELL  9.3 286
	2019 04 02 15 00  1.7  1.7 16.0  0.2  3.4 WNW NNE      SWELL  9.4 288
	2019 04 02 14 00  1.5  1.5 13.8  0.2  3.6 WNW NNW      SWELL  9.1 285
	2019 04 02 13 00  1.7  1.7 17.4  0.2  3.8 WNW NNE      SWELL 10.1 288
	2019 04 02 12 00  1.9  1.8 16.0  0.2  3.7 WNW NNW      SWELL 10.5 289
`

func TestParseString(t *testing.T) {
	surf := string(testString)
	//to match all leading/trailing whitespac
	leadSpace := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	//to match 2 or more whitespace symbols inside a string
	extraSpace := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	final := leadSpace.ReplaceAllString(surf, "")
	final = extraSpace.ReplaceAllString(final, " ")
	rows := strings.Split(final, "\n")
	var dataType []string

	if reflect.TypeOf(rows) !=  reflect.TypeOf(dataType) {
		t.Errorf("Expected %T, got %T", dataType, rows)
	}
}


func TestPerimeter(t *testing.T) {
    rectangle := Rectangle{10.0, 10.0}
    got := Perimeter(rectangle)
    want := 40.0

    if got != want {
        t.Errorf("got %.2f want %.2f", got, want)
    }
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got , want)
		}
	}


	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{ 12, 6}
		checkArea(t , rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}
//checkArea is a helper function which is asking for a Shape to be passed
// The code will not compile unless a shape is being passed, to do this we will need to use an interface declaration "type Shape interface { Area() float 64}"
//Errorf function lets us use formatting features to create descriptive error messages.
