package main

import (
	"fmt"
	"strings"
)

func getUrl() string {
	var url string
	fmt.Print("Enter URL: ")
	fmt.Scan(&url)
	return url
}

func splitUrl(url string) []string {
	urlSlice := strings.Split(url, ".")
	return urlSlice
}

func getTld(urlSlice []string) string {
	var tld string
	tld = urlSlice[len(urlSlice)-1]
	return tld
}

func getSubdomains(urlSlice []string) []string {
	urlSliceEntries := len(urlSlice)
	subdomainSlice := urlSlice[1 : urlSliceEntries-1]
	return subdomainSlice
}

func getHostname(urlSlice []string) string {
	var hostname string
	hostname = urlSlice[0]
	return hostname
}

func main() {
	receivedUrl := getUrl()
	fmt.Println("URL to explode:", receivedUrl)
	urlSlice := splitUrl(receivedUrl)

	tld := getTld(urlSlice)
	hostname := getHostname(urlSlice)

	fmt.Println("Top level domain:", tld)
	fmt.Println("Hostname:", hostname)

	fmt.Println(getSubdomains(urlSlice))
	fmt.Println(getHostname(urlSlice))
}
