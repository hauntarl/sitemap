# Sitemap Builder

One way these can be built is by first visiting the root page of the website and making a list of every link on that page that goes to a page on the same domain. For instance, on calhoun.io you might find a link to calhoun.io/hire-me/ along with several other links.

Once you have created the list of links, you could then visit each and add any new links to your list. By repeating this step over and over you would eventually visit every page that on the domain that can be reached by following links from the root page.

Implementation of Sitemap Builder from gophercises, including the bonus section.

**[Gophercises](https://courses.calhoun.io/courses/cor_gophercises)**  by Jon Calhoun

**Run Commands:**

- go run main.go --help (-h)
- go run main.go
- go run main.go --depth ```int``` (-depth=```int```)

**Features:**

- performing http.Get request for given urls and parsing the documents
- using bread-first traversal on child hrefs to generate a sitemap
- encoding generated sitemap into xml format

**Packages explored:**

- encoding/xml - to encode go data structure into xml format
- flag - to get depth of search and root url
- net/http - to perform GET request on urls
- net/url - to segregate raw urls into accessible data structure
- os - to create new file and store encoded results into
- strings - to check for prefixes in the parsed links
- [github.com/hauntarl/link-parser](github.com/hauntarl/link-parser) - to parse the HTML document and extract all hrefs from it

**Output:**

``` terminal
D:\gophercises\sitemap>go run main.go --help
Usage of C:\Users\hauntarl\AppData\Local\Temp\go-build2610664103\b001\exe\main.exe:
  -depth int
        the maximum depth of links to follow when building a sitemap (default -1)
  -url string
        the url that you want to build a sitemap for (default "https://gophercises.com")

D:\gophercises\sitemap>go run main.go -depth=2
gophercises.com.xml:
<urlset>
  <url>
    <loc>https://gophercises.com/demos/cyoa/denver</loc>
  </url>
  <url>
    <loc>https://gophercises.com</loc>
  </url>
  <url>
    <loc>https://gophercises.com/</loc>
  </url>
  <url>
    <loc>https://gophercises.com/demos/cyoa/</loc>
  </url>
  <url>
    <loc>https://gophercises.com/demos/cyoa/new-york</loc>
  </url>
</urlset>
```
