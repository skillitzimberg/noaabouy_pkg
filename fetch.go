package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"io/ioutil"
	"log"
	"net/http"
)

type Fetch struct {
}

func (f Fetch) Name() string {
	return "fetch"
}

func (f Fetch) Version() string {
	return "1.0"
}

type FetchInput struct {
	URL  string
	Bouy string
}

type FetchOutput struct {
	Data string
}

func (f Fetch) Execute(in step.Context) (interface{}, error) {

	input := FetchInput{URL: "https://www.ndbc.noaa.gov/data/realtime2/", Bouy: "46029.spec"}

	step.BindInputs(&input)

	output := FetchOutput{}

	output.Data = f.GetBouyData(input)
	println("Hello World")
	return nil, nil
}

func (f Fetch) GetBouyData(fetchInput FetchInput) string {

	response, err := http.Get(fetchInput.URL + fetchInput.Bouy)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(responseData)
}