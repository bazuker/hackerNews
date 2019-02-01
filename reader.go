package main

import (
	"fmt"
	"github.com/ungerik/go-rss"
	"os"
	"strconv"
)

const RssFilename = "hacker_news.csv"

func ReadRssFromUrl(rssUrl string, csv bool) {
	// making a http request to get the RSS feed in XML
	// then the package parses it and returns the channel
	channel, err := rss.Read(rssUrl)
	checkErr(err)
	// save everything to the csv file
	if csv {
		// create an output file
		f, err := os.Create(RssFilename)
		checkErr(err)
		// write column headers
		_, err = f.WriteString("\"Number\",\"Title\",\"Link\"\n")
		checkErr(err)
		// flushing the channel item
		for i, item := range channel.Item {
			_, err = f.WriteString("\"" + strconv.Itoa(i+1) + "\",\"" + item.Title + "\",\"" + item.Link + "\"\n")
			checkErr(err)
		}
		// closing up the file
		err = f.Close()
		checkErr(err)
		// letting the user know that the data is saved
		fmt.Println(RssFilename)
		return
	}
	// otherwise just print it
	fmt.Println(channel.Title)
	for i, item := range channel.Item {
		fmt.Println(i+1, item.Title)
		fmt.Println(item.Link)
		if i == 19 { // only print top 20
			break
		}
	}
}
