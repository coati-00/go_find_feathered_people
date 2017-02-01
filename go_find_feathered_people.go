package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/robertkrimen/otto"
)

func findBirdies() {

	// try adopt a pet
	seconddoc, err := goquery.NewDocument("http://www.adoptapet.com/pet-adoption/search/bird/50/miles/10016")
	check(err)
	// first listing doesn't show all details - follow link?
	seconddoc.Find("div.pet_results.rounded_corner").Each(
		func(index int, item *goquery.Selection) {
			mainbird := item.Find("p a.name")
			birdname := mainbird.Text()
			fmt.Printf("Bird Name:  %s\n", birdname)
			detailslink, _ := mainbird.Attr("href")
			followLink(detailslink)
			fmt.Printf("Details Link : %s\n", detailslink)
			birdgender := item.Find(":first-child a").Text()
			fmt.Printf("Bird Gender:  %s\n", birdgender)
			rescueorgtownstate := item.Find("p a.name").Text()
			fmt.Printf("Location: %s\n", rescueorgtownstate)
			phototag := item.Find("span.featured-thumbnail a img")
			photo, _ := phototag.Attr("src")
			fmt.Printf("Photo : %s\n", photo)
		})

}

func main() {
	findBirdies()
}

func getValueFromObject(val otto.Value, key string) (*otto.Value, error) {
	if !val.IsObject() {
		return nil, errors.New("passed val is not an Object")
	}

	valObj := val.Object()

	obj, err := valObj.Get(key)
	if err != nil {
		return nil, err
	}

	return &obj, nil
}

// fatal if there is an error
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func followLink(url string) {

	detailsdoc, err := goquery.NewDocument(url)
	check(err)

	// first listing doesn't show all details - follow link to get them
	detailsdoc.Find("div.container div.row").Each(
		func(index int, item *goquery.Selection) {

			// get bird!
			birdinfobox := item.Find("div.blue_highlight.no_margin.top_margin_xlarge ul li")
			birdinfobox.Each(func(index int, item *goquery.Selection) {
				if item.Find("b").Text() == "Breed" {
					fmt.Printf("Species: %s\n", item.Text())
				}
				if item.Find("b").Text() == "Color:" {
					fmt.Printf("Color: %s\n", item.Text())
				}
				if item.Find("b").Text() == "Age:" {
					fmt.Printf("Age: %s\n", item.Text())
				}
				if item.Find("b").Text() == "Sex:" {
					fmt.Printf("Sex: %s\n", item.Text())
				}
			})

			// birdy info
			item.Find("div.info_box.row div div div.body").Each(
				func(index int, item *goquery.Selection) {
					fmt.Printf("Birdy Info: %s\n", item.Text())
				})

			// birdy large pic  div.col-sm-7 div.large_image img
			item.Find("div.col-sm-7 div.large_image").Each(
				func(index int, item *goquery.Selection) {
					largephoto, _ := item.Find("img").Attr("src")
					fmt.Printf("Large Bird Photo: %s\n", largephoto)
				})

			// use sidebar for rescue info
			rescuesidebar := item.Find("div.body.contact_sidebar.hidden-xs")

			// Go over each of the listed items in the side bar and extract correct info
			rescuesidebar.Find("ul li").Each(func(index int, item *goquery.Selection) {
				if item.Find("b").Text() == "Rescue Group:" {
					fmt.Printf("Each Rescue Name: %s\n", item.Find("a").Text())
				}
				if item.Find("b").Text() == "Phone:" {
					fmt.Printf("Loop Each Rescue Phone: %s\n", item.Find("a").Text())
				}
				if item.Find("b").Text() == "E-mail:" {
					fmt.Printf("Loop Each Rescue E-mail: %s\n", item.Find("a").Text())
				}
				if item.Find("b").Text() == "Website:" {
					fmt.Printf("Loop Each Rescue Website: %s\n", item.Find("a").Text())
				}
				if item.Find("b").Text() == "Address:" {
					fmt.Printf("Loop Each Rescue Address: %s\n", item.Text())
				}
			})

		})

}
