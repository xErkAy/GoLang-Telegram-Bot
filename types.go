package main

type Chat struct {
	ID         int64
	messageID  int
	user_id    string
	username   string
	language   string
	first_name string
}

type User struct {
	user_id    string
	username   string
	language   string
	first_name string
}

type Anagram struct {
	Result []string `json:"result"`
}

type Location struct {
	Location *Address `json:"address"`
}

type Address struct {
	Country string `json:"CntryName"`
	City    string `json:"City"`
	Address string `json:"Address"`
}
