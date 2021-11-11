package scraper

import (
	"github.com/gocolly/colly"
	"github.com/jimsmart/grobotstxt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func SiteMapCheck(URL string) (int, string){
	// Array containing all the known URLs in a sitemap
	knownUrls := []string{}
	// Create a Collector specifically for Shopify
	c := colly.NewCollector()

	// Create a callback on the XPath query searching for the URLs
	c.OnXML("//urlset/url/loc", func(e *colly.XMLElement) {
		knownUrls = append(knownUrls, e.Text)
	})

	// Start the collector
	c.Visit(URL + "/sitemap.xml")

	/*log.Println("All known URLs:")
	  for _, url := range knownUrls {
	  	log.Println("\t", url)
	  }*/

	if len(knownUrls) > 0 {
		return 0, "Sitemap successfully found for "+ URL
	}

	if len(knownUrls) == 0 {
		return 2, "Sitemap found but No URLs exist for "+ URL+". Kindly check if it is in the conventional format as defined at https://www.sitemaps.org/protocol.html"
	}
	return 1,"Sitemap doesn't exist!"
}

func RobotsTXT(URL string) (int,string) {

	resp, err := http.Get(URL+"/robots.txt")
	if err != nil {
		log.Fatalln(err)
	}
	//Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		// Return 1 indicating no robots.txt found
		return 1,"Robots.txt doesnt exist"
	}
	//Convert the body to type string
	robotsTxt := string(body)

	// Check if Sitemap exists
	sitemaps := grobotstxt.Sitemaps(robotsTxt)
	if len(sitemaps)>0 {
		log.Println(len(sitemaps))
		// Successfully found sitemap
		return 0,"Successfully found Sitemaps!"
	}
	// Robots.txt found but no sitemaps found
	return 2,"No Sitemaps found!"
}

