package main

import (
	"fmt"
)

type ShortenUrl struct {
	Url string `json:"url"`
}

func (s ShortenUrl) PrintName() {
	fmt.Println(s.Url)
}
