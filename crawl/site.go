package crawl

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/html"
)

// TODO Get any uris - img, script, css uris

// map urls to query
// TODO normalize urls
var cache = make(map[string]*Query)

// urls to parse
var queue = make(chan string)

// anchors gets all the anchor hrefs from the documents
func anchors(doc *html.HtmlDocument) []string {
	// TODO What to do about internal parsing errors?
	nodes, _ := doc.Search("//a/@href")
	hrefs := make([]string, len(nodes))
	for i, node := range nodes {
		hrefs[i] = node.String()
	}
	return hrefs
}

func query(u string) {
	// If the query was already performed and was non-200, return that
}

func Site(rawURL string, threads int) error {
	// Confirm the url is valid
	u, err := normalize(rawURL)
	if err != nil {
		return err
	}

	// TODO Save the scheme and host for links

	// Perform the first query - no From because it is root
	q := &Query{From: "-"}
	if q.Request, err = http.NewRequest("GET", u.String(), nil); err != nil {
		return err
	}

	// cache the query
	cache[u.String()] = q

	var client http.Client
	start := time.Now()
	if q.Response, err = client.Do(q.Request); err != nil {
		return err
	}
	q.Latency = time.Now().Sub(start)

	log.Println(q)

	page, err := ioutil.ReadAll(q.Response.Body)
	if err != nil {
		return err
	}
	doc, err := gokogiri.ParseHtml(page)
	if err != nil {
		return err
	}
	defer doc.Free()
	hrefs := anchors(doc)
	log.Println(hrefs)
	return nil
}
