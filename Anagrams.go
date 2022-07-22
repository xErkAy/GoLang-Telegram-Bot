package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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
