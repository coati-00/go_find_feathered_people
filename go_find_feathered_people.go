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

/* SAMPLE LISTING - Adopt A Pet

<div class="pet_results rounded_corner ">

  <span class="featured-thumbnail">
    <a class="smaller_line_height track" href="http://www.adoptapet.com/pet/16761951-woodbridge-new-jersey-bird" target="_top" data-track='{"eventCat": "Pet Details", "eventAct": "Search Results", "eventLbl": "Click on pet 3"}'>
      <img src="https://s3.amazonaws.com/pet-uploads.adoptapet.com/3/d/7/231622278.jpg" width="200" alt="Adopt A Pet :: Jakers - Woodbridge, NJ" height="200" border="0" />
    </a>
  </span>

  <p class="truncate no_top_margin">
    <a class="name museo700 track" href="http://www.adoptapet.com/pet/16761951-woodbridge-new-jersey-bird" target="_top" data-track='{"eventCat": "Pet Details", "eventAct": "Search Results", "eventLbl": "Click on pet 3"}'>
		Jakers
    </a>
  </p>

  <p class="truncate no_top_margin" style="margin-bottom:0;">
    <a class="smaller_line_height track"
    	 href="http://www.adoptapet.com/pet/16761951-woodbridge-new-jersey-bird" target="_top" data-track='{"eventCat": "Pet Details", "eventAct": "Search Results", "eventLbl": "Click on pet 3"}'>
  		Woodbridge, NJ
    </a>
  </p>

  <a class="smaller_line_height track"
  	 href="http://www.adoptapet.com/pet/16761951-woodbridge-new-jersey-bird" target="_top" data-track='{"eventCat": "Pet Details", "eventAct": "Search Results", "eventLbl": "Click on pet 3"}'>
    Male
  </a>

</div>

*/
/*   SAMPLE LISTING ELEMENT - Pet Finder

<li>
<div class="adoptablePets-item">
   <figure>
     <a href="/petdetail/36448558" data-category="pet details" data-action="click to pet details page">
        <img src="https://drpem3xzef3kf.cloudfront.net/photos/pets/36448558/1/?bust=1476200857&amp;width=186&amp;no_scale_up=1" alt="Demetria and Billy Boy - Parakeet (Other)">
     </a>
    </figure>


    <div class="pet-name-container">
      <h2>
         <a href="/petdetail/36448558" data-category="pet details" data-action="click to pet details page">Demetria and Billy Boy</a>
      </h2>
    </div>
    <p class="breed">Parakeet (Other)</p>
    <p class="specs">Adult • Female • Small</p>
    <p class="rescue-info">
      <a href="http://www.petfinder.com/shelters/NJ775.html">Lonely Grey Rescue</a>
      <br>
      <span>Woodbridge</span>
      <span>NJ</span>
    </p>
</div>
</li>


*/
