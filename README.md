# Sitemap Builder

One way these can be built is by first visiting the root page of the website and making a list of every link on that page that goes to a page on the same domain. For instance, on calhoun.io you might find a link to calhoun.io/hire-me/ along with several other links.

Once you have created the list of links, you could then visit each and add any new links to your list. By repeating this step over and over you would eventually visit every page that on the domain that can be reached by following links from the root page.

Implementation of Sitemap Builder from gophercises, including the bonus section.

**[Gophercises](https://courses.calhoun.io/courses/cor_gophercises)**  by Jon Calhoun

<!--
**Run Commands:**

- go run examples\ex1\main.go
- go run examples\ex2\main.go
- go run examples\ex3\main.go
- go run examples\ex4\main.go

**Features:**

- creating io.Reader from string type
- using depth first search to traverse html document
- extracting relevant information from html document

**Packages explored:**

- fmt
- strings - to create io.Reader and format relevant data
- [golang.org/x/net/html](golang.org/x/net/html) - to parse the html document into Tree structure

**Output:**

``` terminal
D:\gophercises\link-parser>go run examples\ex1\main.go
/other-page                                       : A link to another page some span
/page-two                                         : A link to second page

D:\gophercises\link-parser>go run examples\ex2\main.go
https://www.twitter.com/joncalhoun                : Check me out on twitter
https://github.com/gophercises                    : Gophercises is on Github!

D:\gophercises\link-parser>go run examples\ex3\main.go
#                                                 : Login
/lost                                             : Lost? Need help?
https://twitter.com/marcusolsson                  : @marcusolsson

D:\gophercises\link-parser>go run examples\ex4\main.go
/dog-cat                                          : dog cat
```
