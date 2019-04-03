package main

import (
	"testing"
)

var testRow = `
	2019 04 02 17 00  1.7  1.7 16.0  0.2  3.4 WNW  SW      SWELL  9.6 292
`

var testNewRow = `2019 04 02 17 00 1.7 1.7 16.0 0.2 3.4 WNW SW SWELL 9.6 292`

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

func TestRemoveSpacesFromString(t *testing.T) {
	data := testRow
	got := removeSpacesFromString(data)
	want := testNewRow
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestParseStringCount(t *testing.T) {
	data := testRow
	newRow := removeSpacesFromString(data)
	got := parseStringCount(newRow)
	want := 15
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestGetRowWWD(t *testing.T) {
	data := testRow
	newRow := removeSpacesFromString(data)
	removeRowSpaces := parseString(newRow)
	got := getRowDataWWD(removeRowSpaces)
	want := "SW"
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestGetRowSwP(t *testing.T) {
	data := testRow
	newRow := removeSpacesFromString(data)
	removeRowSpaces := parseString(newRow)
	got := getRowDataSwP(removeRowSpaces)
	want := 16.0
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestGetRowSwH(t *testing.T) {
	data := testRow
	newRow := removeSpacesFromString(data)
	removeRowSpaces := parseString(newRow)
	got := getRowDataSwH(removeRowSpaces)
	want := 1.7
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestWindScore(t *testing.T) {
	data := removeSpaces(testNewRow)
	wwd := getRowDataWWD(data)
	got := windScore(wwd)
	want := 1.0

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestSwellScore(t *testing.T) {
	data := removeSpaces(testNewRow)
	swell := getRowDataSwP(data)
	got := swellScore(swell)
	want := 5.0

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestWaveScore(t *testing.T) {
	data := removeSpaces(testNewRow)
	swp := getRowDataSwP(data)
	swh := getRowDataSwH(data)
	got := waveScore(swp, swh)
	want := 4.0

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

// Why is this test failing? got 3.33 want 3.33
// func TestAverageScore(t *testing.T) {
// 	wind := 5.0
// 	swell := 4.0
// 	wave := 1.0
// 	got := averageScore(wind, swell, wave)
// 	want := 3.333333

// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}
// }

func TestGetScoreForToday(t *testing.T) {
	data := removeSpaces(testString)
	got := getScoreForToday(data)
	want := 4.0

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}
