package weather

import (
	"fmt"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"strings"
)

type Weather struct {
	City string `json:"city"`
	Time string `json:"time"`
	Temp string `json:"temp"`
}

func (w Weather) Show() {
	fmt.Printf("%+v", w)
}

func ParceWeather()  {
	parcer := geziyor.NewGeziyor(&geziyor.Options{
		StartURLs:   []string{"https://www.gismeteo.ua/weather-kharkiv-5053/now/"},
		ParseFunc:   parse,
		LogDisabled: true,
	})
	parcer.Start()
}

func parse(g *geziyor.Geziyor, r *client.Response) {
	w := Weather{}

	doc := r.HTMLDoc
	find := doc.Find("#js-search")
	if attr, exists := find.Attr("placeholder"); exists {
		w.City = attr
	}

	time := doc.Find("#time")
	first := time.Children().First()
	w.Time = first.Text()

	temp := doc.Find(".nowvalue__text_l")
	w.Temp = strings.TrimSpace(temp.First().Text())

	w.Show()
}