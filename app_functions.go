package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getLocation(latitude float64, longitude float64) string {

	url := fmt.Sprintf("https://geocode.arcgis.com/arcgis/rest/services/World/GeocodeServer/reverseGeocode?outSR=4326&returnIntersection=false&location=%f,%f&f=json", longitude, latitude)
	response, err := http.Get(url)
	checkErrors(err)

	body, err := ioutil.ReadAll(response.Body)
	checkErrors(err)

	var responseObject Location
	json.Unmarshal(body, &responseObject)

	return fmt.Sprintf("%s, %s, %s", responseObject.Location.Country, responseObject.Location.City, responseObject.Location.Address)
}

func solveAnagram(word string) string {

	url := fmt.Sprintf("https://anagram.poncy.ru/anagram-decoding.cgi?name=anagram_main&inword=%s&answer_type=1", word)
	response, err := http.Get(url)
	checkErrors(err)

	body, err := ioutil.ReadAll(response.Body)
	checkErrors(err)

	var responseObject Anagram
	json.Unmarshal(body, &responseObject)

	result := "Результаты решения анаграммы:\n"

	for i, s := range responseObject.Result {
		result += fmt.Sprintf("%d. %s\n", i+1, s)
	}

	return result
}

func makeAnagram(word string) string {

	url := fmt.Sprintf("https://anagram.poncy.ru/anagram-decoding.cgi?name=anagram_main&inword=%s&answer_type=2", word)
	response, err := http.Get(url)
	checkErrors(err)

	body, err := ioutil.ReadAll(response.Body)
	checkErrors(err)

	var responseObject Anagram
	json.Unmarshal(body, &responseObject)

	return "Составленная анаграмма: " + responseObject.Result[0]
}
