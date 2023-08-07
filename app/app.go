package app

import (
	"log"
	"net/url"
	"os"
)

var RootURL *url.URL

func AbsoluteURL(u url.URL) url.URL {
	u.Scheme = RootURL.Scheme
	u.Host = RootURL.Host
	return u
}

func init() {
	var err error
	root := os.Getenv("ROOT_URL")
	if root == "" {
		panic("ROOT_URL is not set")
	}
	RootURL, err = url.Parse(root)
	if err != nil {
		panic(err)
	}
	log.Printf("RootURL = %s", RootURL.String())
}
