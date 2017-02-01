package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/robertkrimen/otto"
)

func findBirdies() {
	// petfinder has dynamic pages will deal with them later

	/*	// get search for local area
		doc, err := goquery.NewDocument("https://www.petfinder.com/pet-search?location=10001&animal=bird&breed=&distance=25")
		check(err)
		doc.Find("li div").Each(
			func(index int, item *goquery.Selection) {
				//fmt.Printf("Item: %s\n", item)
				//name := item.Find("div .pet-name-container").Text()
				//name := item.Find(".pet-name-container").Text()
				//fmt.Printf("Name of birdy: %s\n", name)
				rescueinfo := item.Find("h2 a").Text()
				fmt.Printf("Rescue Info: %s\n", rescueinfo)
			})*/

	// most of page seems to be generated dynamically with angular
	// and mustache templates --> see if otto can get js values

	//find the JavaScript
	// script := doc.Find("script").First()
	// fmt.Printf("First Script: %s\n\n\n", script.Text())

	// vm := otto.New()

	// // to avoid errors prior to returned info
	// _, err = vm.Run("var window = {}")
	// check(err)

	// // not sure if this is evaluating or running the script
	// // "scrry scrry"
	// _, err = vm.Run(script.Text())
	// check(err)

	// // traverse down the object path: window > POST_DATA > <videoID> > videoDashURL
	// wVal, err := vm.Get("window")
	// check(err)

	// pdata, err := getValueFromObject(wVal, "POST_DATA")
	// check(err)

	// if pdata != nil {
	// 	fmt.Printf("pdata: %s\n\n\n", pdata)
	// }
	//videoData, err := getValueFromObject(*pdata, id)

	//videoDashUrl, err := getValueFromObject(*videoData, "videoDashUrl")

	// select elements wanted
	/*doc.Find("div.adoptablePets-item").Each(
	func(index int, item *goquery.Selection) {
		//fmt.Printf("Item: %s\n", item)
		//name := item.Find("div .pet-name-container").Text()
		name := item.Find(".pet-name-container")
		fmt.Printf("Name of birdy: %s\n", name)
		/*birdyspecs =
		species =
		rescueorgname =
		rescueorgtown =
		rescueorgstate = */
	//		})

	/*	// try another way to get pet names
		sel := doc.Find("div.pet-name-container")
		for element := range sel.Nodes {
			single := sel.Eq(element)
			fmt.Printf("Single: %s\n", single)
		}*/

	// go over all scripts
	/*	doc.Find("script").Each(
		func(index int, item *goquery.Selection) {
			scripttext := item.Text()
			fmt.Printf("Rescue Info: %s\n", scripttext)
		})*/

	// try adopt a pet
	seconddoc, err := goquery.NewDocument("http://www.adoptapet.com/pet-adoption/search/bird/50/miles/10016")
	check(err)
	//
	seconddoc.Find("div.pet_results.rounded_corner").Each(
		func(index int, item *goquery.Selection) {
			//fmt.Printf("Item: %s\n", item)
			//name := item.Find("div .pet-name-container").Text()
			//name := item.Find(".pet-name-container").Text()
			//fmt.Printf("Name of birdy: %s\n", name)
			birdname := item.Find("p a.name").Text()
			fmt.Printf("Bird Name:  %s \n", birdname)
			birdgender := item.Find("a").Text()
			fmt.Printf("Bird Gender:  %s  \n", birdgender)
			//species := item.Find("p a.name").Text()
			//rescueorgname := item.Find("p a.name").Text()
			//rescueorg := item.Find("p a").Text()
			rescueorgtownstate := item.Find("p a.name").Text()
			fmt.Printf("Location:   %s  \n", rescueorgtownstate)
			//photo := item.Find("span.featured-thumbnail a img").Attr('src')
		})

}

func main() {
	findBirdies()
}

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
