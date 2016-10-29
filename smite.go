package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Smite struct {
	Gods []*gods
}

type gods struct {
	Name       string
	Pantheon   string
	Attack     string
	Power      string
	Class      string
	Difficulty string
}

func (smite *Smite) UpdateGods() {
	doc, err := goquery.NewDocument("http://smite.gamepedia.com/List_of_gods")
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("table tbody tr").Each(func(i int, s *goquery.Selection) {
		td := s.Find("td")
		if td.Text() == "" {
			return
		}
		smite.Gods = append(smite.Gods, &gods{
			Name:     td.Eq(1).Text(),
			Pantheon: strings.Replace(td.Eq(2).Text(), " ", "", 1),
			Attack:   td.Eq(3).Text(),
			Power:    strings.Replace(td.Eq(4).Text(), " ", "", 1),
			Class:    strings.Replace(td.Eq(5).Text(), " ", "", 1),
		})
	})

	for _, god := range smite.Gods {
		doc, err := goquery.NewDocument("http://smite.gamepedia.com/" + god.Name)
		if err != nil {
			return
		}

		god.Difficulty = doc.Find("table.infobox tr").Eq(8).Find("td").Text()
	}
}
