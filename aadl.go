package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/publicsuffix"
)

func main() {
	username := flag.String("username", "", "AADL username")
	password := flag.String("password", "", "AADL password")
	flag.Parse()
	if *username == "" || *password == "" {
		flag.Usage()
		return
	}

	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	client := &http.Client{Jar: jar}

	// get the form_id and form_build_id, whatever the heck those are.
	req, _ := http.NewRequest("GET", "https://aadl.org/user/login", nil)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	formID, _ := doc.Find("input[name=form_id]").Attr("value")
	formBuildID, _ := doc.Find("input[name=form_build_id]").Attr("value")

	req, _ = http.NewRequest("POST", "https://aadl.org/user/login",
		strings.NewReader(url.Values{
			"name":          {*username},
			"pass":          {*password},
			"form_id":       {formID},
			"form_build_id": {formBuildID},
		}.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode >= 400 {
		panic(resp.StatusCode)
	}

	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	doc.Find("tr.checkout-row").Each(func(i int, s *goquery.Selection) {
		cols := []string{}
		s.Find("td").Each(func(_ int, td *goquery.Selection) {
			cols = append(cols, td.Text())
		})
		fmt.Println(cols[1], cols[2], cols[4])
	})
}
