// oi
package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
)

//Result exported
type Result struct {
	userName string
	title    string
	likes    string
}

func (r Result) String() string {
	return fmt.Sprint(r.userName, " - ", r.title, " - ", r.likes, " claps")
}

func main() {
	urlToProcess := []string{
		"https://medium.freecodecamp.org/how-to-columnize-your-code-to-improve-readability-f1364e2e77ba",
		"https://medium.freecodecamp.org/how-to-think-like-a-programmer-lessons-in-problem-solving-d1d8bf1de7d2",
		"https://medium.freecodecamp.org/code-comments-the-good-the-bad-and-the-ugly-be9cc65fbf83",
		"https://uxdesign.cc/learning-to-code-or-sort-of-will-make-you-a-better-product-designer-e76165bdfc2d",
	}
	var ini time.Time
	fmt.Println("Without go routine:")
	ini = time.Now()
	for _, url := range urlToProcess {
		r := scrap(url)
		fmt.Println(r)
	}
	fmt.Println("(Took ", time.Since(ini).Seconds(), "secs)")

}

func hasClass(attribs []html.Attribute, className string) bool {
	for _, attr := range attribs {
		if attr.Key == "class" && strings.Contains(attr.Val, className) {
			return true
		}
	}
	return false
}

func getFirstTextNode(htmlParsed *html.Node) *html.Node {
	if htmlParsed == nil {
		return nil
	}

	for m := htmlParsed.FirstChild; m != nil; m = m.NextSibling {
		if m.Type == html.TextNode {
			return m
		}
		r := getFirstTextNode(m)
		if r != nil {
			return r
		}
	}
	return nil
}

func getFirstElementByClass(htmlParsed *html.Node, elm, className string) *html.Node {
	for m := htmlParsed.FirstChild; m != nil; m = m.NextSibling {
		if m.Data == elm && hasClass(m.Attr, className) {
			return m
		}
		r := getFirstElementByClass(m, elm, className)
		if r != nil {
			return r
		}
	}
	return nil
}

func scrap(url string) (r Result) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("ERROR: It can't scrap '", url, "'")
	}
	// Close body when function ends
	defer resp.Body.Close()
	body := resp.Body
	htmlParsed, err := html.Parse(body)
	if err != nil {
		fmt.Println("ERROR: It can't parse html '", url, "'")
	}

	a := getFirstTextNode(getFirstElementByClass(htmlParsed, "a", "ds-link--styleSubtle"))
	if a != nil {
		r.userName = a.Data
	} else {
		fmt.Println("Scrap error: Can't find username. url:'", url, "'")
	}

	div := getFirstElementByClass(htmlParsed, "div", "section-content")
	h1 := getFirstTextNode(getFirstElementByClass(div, "h1", "graf--title"))
	if h1 != nil {
		r.title = h1.Data
	} else {
		fmt.Println("Scrap error: Can't find title. url:'", url, "'")
	}

	footer := getFirstElementByClass(htmlParsed, "footer", "u-paddingTop10")
	buttonLikes := getFirstTextNode(getFirstElementByClass(footer, "button", "js-multirecommendCountButton"))
	if buttonLikes != nil {
		r.likes = buttonLikes.Data
	} else {
		fmt.Println("Scrap error: Can't find button of likes. url:'", url, "'")
	}

	return
}
