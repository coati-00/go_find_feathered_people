package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func postScrape() {

	doc, err := goquery.NewDocument("https://www.petfinder.com/pet-search?location=10001&animal=bird&breed=&distance=100")
	if err != nil {
		log.Fatal(err)
	}

	// select elements wanted
	doc.Find("li,div.adoptablePets-item").Each(func(index int, item *goquery.Selection) {
		name := item.Find("div.pet-name-container h2 a ").Text()
		fmt.Printf("Name of birdy: %s\n", name)
		/*birdyspecs =
		species =
		rescueorgname =
		rescueorgtown =
		rescueorgstate = */

	})

}

/*   SAMPLE LISTING ELEMENT

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
