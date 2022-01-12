package scraper

import (
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func MetaTags(w http.ResponseWriter, r *http.Request){
	URL:=r.URL.Query().Get("url")
	//Verify if URL exists
	res, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("head").Each(func(i int, s *goquery.Selection) {
		n := s.Find("meta").Nodes
		for _,v:= range n{
			x:=v.Attr
			for _,y:= range x{
				log.Println(y.Val)
			}
		}
	})
}

