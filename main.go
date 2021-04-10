package main

import (
	"encoding/xml"
	"flag"
	"io"
	"math"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/hauntarl/link-parser"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

var (
	root  *string
	depth *int
)

func init() {
	root = flag.String(
		"url", "https://gophercises.com",
		"the url that you want to build a sitemap for",
	)
	depth = flag.Int(
		"depth", math.MaxInt8,
		"the maximum depth of links to follow when building a sitemap",
	)
	flag.Parse()
}

type (
	loc struct {
		Value string `xml:"loc"`
	}
	urlset struct {
		Urls  []loc  `xml:"url"`
		Xmlns string `xml:"xmlns,attr"`
	}
)

func main() {
	var (
		seen = map[string]struct{}{}
		que  = []string{*root}
	)
	for *depth >= 0 && len(que) != 0 {
		que = traverse(que, seen)
		*depth--
	}

	data := urlset{make([]loc, 0, len(seen)), xmlns}
	for page := range seen {
		data.Urls = append(data.Urls, loc{page})
	}

	file, err := create()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString(xml.Header)
	enc := xml.NewEncoder(file)
	enc.Indent("", "  ")
	if err := enc.Encode(data); err != nil {
		panic(err)
	}
}

// traverse iterates overs every page present in queue:
//
// 1. check if the link is already visited
//
// 2. if not, get all the child links for cur page
//
// 3. accumulate all children into next queue and return it
func traverse(que []string, seen map[string]struct{}) (nex []string) {
	for _, page := range que {
		if _, ok := seen[page]; ok {
			continue
		}
		seen[page] = struct{}{}
		nex = append(nex, get(page)...)
	}
	return
}

// get makes an http.Get request for a given page:
//
// 1. parse the HTML to extract all the hrefs
//
// 2. filter and keep the links which belong to Request.URL domain
//
// 3. return the result
func get(page string) []string {
	resp, err := http.Get(page)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// io.Copy(os.Stdout, resp.Body) // print the content onto terminal

	reqURL := &url.URL{
		Scheme: resp.Request.URL.Scheme,
		Host:   resp.Request.URL.Host,
	}
	baseURL := reqURL.String()
	return filter(hrefs(resp.Body, baseURL), withPrefix(baseURL))
}

// hrefs parses the HTML and transforms the path into appropriate url
func hrefs(r io.Reader, baseURL string) (res []string) {
	links, _ := link.Parse(r)
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			res = append(res, baseURL+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			res = append(res, l.Href)
		}
	}
	return
}

type predicate func(string) bool

func filter(pages []string, apply predicate) (res []string) {
	for _, p := range pages {
		if apply(p) {
			res = append(res, p)
		}
	}
	return
}

func withPrefix(val string) predicate {
	return func(page string) bool {
		return strings.HasPrefix(page, val)
	}
}

// create provides a new io.Writer for xml.Encoder to write data into:
//
// 1. extract the url.Host from given root
//
// 2. create a new file with the extracted host name
//
// 3. return the file pointer
func create() (*os.File, error) {
	baseURL, err := url.Parse(*root)
	if err != nil {
		return nil, err
	}
	return os.Create(baseURL.Host + ".xml")
}
