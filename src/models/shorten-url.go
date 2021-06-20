package models

import (
	"fmt"
	"github.com/dchest/uniuri"
	"gorm.io/gorm"
)

type ShortenUrl struct {
	gorm.Model
	Url string `gorm:"type:varchar(255); not null" json:"url"`
	HashedUrl string `gorm:"type:varchar(50); unique_index; not null" json:"hashedUrl"`
	VisitCount int `gorm:"type:integer; default '0'" json:"visitCount"`
}

func (s ShortenUrl) PrintName() {
	fmt.Println(s.Url)
}

func (s ShortenUrl) GetShortenUrl() string {
	return uniuri.New()
}
