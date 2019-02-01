package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestReadRssFromUrl(t *testing.T) {
	const testRssUrl = "https://news.ycombinator.com/rss"
	ReadRssFromUrl(testRssUrl, true)
	if _, err := os.Stat(RssFilename); os.IsNotExist(err) {
		t.Error("csv file was not saved")
	}
	// remove the file after the test
	defer os.Remove(RssFilename)
	// open the file
	f, _ := os.Open(RssFilename)
	r := csv.NewReader(bufio.NewReader(f))
	defer f.Close()
	// go over the records
	skipHeaders := false
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		// there must be only three columns
		if len(record) != 3 {
			t.Error("file does not have required columns")
			break
		}
		// skip the first column of headers
		if !skipHeaders {
			skipHeaders = true
			continue
		}
		// validate the number
		number := record[0]
		if _, err := strconv.Atoi(number); err != nil {
			t.Error("first column is not a number")
			break
		}
		// validate the link
		link := record[2]
		if !strings.HasPrefix(link, "http") {
			t.Log("third column is not a link")
		}
	}
}
