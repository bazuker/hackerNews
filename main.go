package main

/*
	Author: Daniil Furmanov
	Description: a program that reads RSS feed from a remote host
		and is capable of saving the results in CSV format
*/

import (
	"os"
	"strings"
)

func main() {
	const rssUrl = "https://news.ycombinator.com/rss"
	// check for -csv argument
	csv := false
	if len(os.Args) > 1 {
		csv = strings.ToLower(os.Args[1]) == "-csv"
		// could have used flag package but it's too much for a single attribute
	}
	// read the feed
	ReadRssFromUrl(rssUrl, csv)
}
