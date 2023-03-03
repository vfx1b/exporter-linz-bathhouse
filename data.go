package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

var (
	apiUrl = "https://www.linzag.at/RESTfulBaths/rest/utilisation/"

	// locations can be regenrated with this script on the website of the city
	// let elements = document.getElementsByClassName("lag-workload-teaser-text");
	// let output = "";
	//
	//	for (let i = 0; i < elements.length; i++) {
	//		let e = elements[i];
	//		let elementIdStr = e.id;
	//		let serverSideId = elementIdStr.split("-")[1];
	//		output += "{ServerSideId: " + parseInt(serverSideId) + " , Name: \"" + e.innerText + "\"},\n"
	//	}
	//
	// console.log(output);
	locations = []Location{
		{ServerSideId: 10, Name: "Parkbad Eisbetrieb"},
		{ServerSideId: 3, Name: "Parkbad Hallenbad"},
		{ServerSideId: 6, Name: "Parkbad Sauna"},
		{ServerSideId: 11, Name: "Parkbad Parken"},
		{ServerSideId: 13, Name: "Hummelhof Hallenbad"},
		{ServerSideId: 17, Name: "Hummelhof Sauna"},
		{ServerSideId: 18, Name: "Hummelhof Wellness"},
		{ServerSideId: 9, Name: "Hummelhof Parken"},
		{ServerSideId: 1, Name: "Biesenfeld Hallenbad"},
		{ServerSideId: 20, Name: "Biesenfeld Sauna"},
		{ServerSideId: 2, Name: "Biesenfeld Parken"},
		{ServerSideId: 14, Name: "Schörgenhub Hallenbad"},
		{ServerSideId: 23, Name: "Schörgenhub Sauna"},
		{ServerSideId: 7, Name: "Schörgenhub Freieis"},
		{ServerSideId: 4, Name: "Ebelsberg Sauna"},
		{ServerSideId: 15, Name: "Ebelsberg Freieis"},
	}
)

func fetchMetrics() *[]Payload {
	payloadResponse := []Payload{}
	for _, location := range locations {
		req, err := http.NewRequest(http.MethodGet, apiUrl+strconv.Itoa(int(location.ServerSideId)), nil)
		if err != nil {
			log.Error(err)
			return nil
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Error(err)
			return nil
		}
		b, err := io.ReadAll(res.Body)
		if err != nil {
			log.Error(err)
			return nil
		}
		payload := []Payload{}
		err = json.Unmarshal(b, &payload)
		if err != nil {
			log.Error(err)
			return nil
		}
		payloadResponse = append(payloadResponse, payload[0])
	}
	return &payloadResponse
}
