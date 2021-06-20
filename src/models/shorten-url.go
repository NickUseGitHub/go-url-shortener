package models

import (
	"fmt"
	"github.com/dchest/uniuri"
	"gorm.io/gorm"
)

type ShortenUrl struct {
	gorm.Model
	Url string `json:"url"`
}

func (s ShortenUrl) PrintName() {
	fmt.Println(s.Url)
}

func (s ShortenUrl) GetShortenUrl() string {
	return uniuri.New()
}
