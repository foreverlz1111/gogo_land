package main

import (
	"encoding/json"
	"log"
)

type Jsonexample struct {
	a  string
	b  string `json:"B"`
	Cs string
	D  string `json:"DD"` // 必须大写，有tag优先tag
}

func main() {
	jsonexample := Jsonexample{
		a:  "11",
		b:  "22",
		Cs: "33",
		D:  "44",
	}
	log.Println(jsonexample)
	jsonformat, _ := json.Marshal(jsonexample)
	log.Printf("%+v", string(jsonformat))
}
