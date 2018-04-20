package scraper

import (
	"log"
	"regexp"
	"strconv"

	"github.com/gocolly/colly"
	"github.com/techmexdev/lineuplist"
)

// Festivals scrapes musicfestivalwizard.com for
// festivals in the US.
func Festivals() ([]lineuplist.Festival, error) {
	fests := []lineuplist.Festival{}

	for i := 1; i <= pageCount(); i++ {
		fests = append(fests, scrapePage(i)...)
	}

	return fests, nil
}

func pageCount() int {
	var lastPageNum int
	c := colly.NewCollector()

	c.OnHTML("ul.page-numbers", func(e *colly.HTMLElement) {
		pageNumEls := e.DOM.Find("a.page-numbers")
		var err error
		lastPageNum, err = strconv.Atoi(pageNumEls.Eq(pageNumEls.Length() - 2).Text())
		if err != nil {
			log.Fatalf("could not parse number of pages: %s", err.Error())
		}
	})

	c.Visit("https://www.musicfestivalwizard.com/all-festivals/?festival_guide=us-festivals")
	c.Wait()

	return lastPageNum
}

func scrapePage(page int) []lineuplist.Festival {
	fests := []lineuplist.Festival{}
	c := colly.NewCollector()

	c.OnHTML(".singlefestlisting", func(e *colly.HTMLElement) {
		name := regexp.MustCompile("(.*[^ \\d])").FindString(e.ChildText(".festivaltitle"))
		startDate, endDate, err := parseDate(e.ChildText(".festivaldate"))
		if err != nil {
			log.Println("error parsing date for: ", name, err)
			return
		}

		country, state, city, err := parseLocation(e.ChildText(".festivallocation"))
		if err != nil {
			log.Println("error parsing location for ", name, err)
			return
		}

		fests = append(fests, lineuplist.Festival{
			Name:      name,
			StartDate: startDate,
			EndDate:   endDate,
			Country:   country,
			State:     state,
			City:      city,
		})
	})

	c.Visit("https://www.musicfestivalwizard.com/all-festivals/page/" + strconv.Itoa(page) + "/?festival_guide=us-festivals")
	c.Wait()

	return fests
}
