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
