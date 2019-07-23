package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	NRKFlow()
}

func NRKFlow() {
	minute := "00"
	hour := "12"
	month := "07"
	year := "19"
	var day string
	for i := 1; i < 31; i++ {
		if i < 10 {
			day = "0" + strconv.Itoa(i)
		} else {
			day = strconv.Itoa(i)
		}
		fileUrl := NRKNewsURL(hour, day, month, year)
		fmt.Println(fileUrl)
		if err := DownloadFile("nrk/"+"20"+year+"-"+month+"-"+day+"-"+hour+minute+".xml", fileUrl); err != nil {
			panic(err)
		}
	}
}

// NRKNewsFileName creates the correct filename for any given date and time of news broadcast
// Eg: https://tv.nrk.no/serie/nyheter/201907/NNFA41013519
// From News @ 21:10 28/07 2019 the string should be: NNFA41013519
// Meanings:
// NNFA
// 410135
// 19 : Year
// From News @ 12:00 22/07 2019 :  NNFA12072219AH
func NRKNewsURL(hour string, day string, month string, year string) (fileurl string) {

	prefix := "NNFA"
	suffix := "AH"
	return "https://tv.nrk.no/programsubtitles/" + prefix + hour + month + day + year + suffix
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
