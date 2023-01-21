package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"testing"

	"fyne.io/fyne/v2/test"
)

var testApp Config

func TestMain(m *testing.M) {
	a := test.NewApp()
	testApp.App = a
	testApp.HTTPClient = client
	os.Exit(m.Run())
}

var jsonToReturn = `
{
	"ts": 1674305796135,
	"tsj": 1674305795317,
	"date": "Jan 21st 2023, 07:56:35 am NY",
	"items": [
		{
		"curr": "USD",
		"xauPrice": 1926,
		"xagPrice": 23.934,
		"chgXau": -5.935,
		"chgXag": 0.0625,
		"pcXau": -0.3072,
		"pcXag": 0.2618,
		"xauClose": 1932.015,
		"xagClose": 23.8715
		}
	]
}`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

var client = NewTestClient(func(req *http.Request) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
		Header:     make(http.Header),
	}
})
