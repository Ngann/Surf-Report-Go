package main

import (
	"testing"
)

var testRow = `
	2019 04 02 17 00  1.7  1.7 16.0  0.2  3.4 WNW  SW      SWELL  9.6 292
`

var testNewRow = `2019 04 02 17 00 1.7 1.7 16.0 0.2 3.4 WNW SW SWELL 9.6 292`

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
