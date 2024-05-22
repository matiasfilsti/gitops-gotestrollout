package main

import (
	"fmt"
	"net/url"
	"os"
	"net/http"
	"time"
)

func main() {

	tr := &http.Transport{
		MaxIdleConns:       1,
		IdleConnTimeout:    10 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	//urlArgs := os.Args[1:]
	urlArgs := []string{"http://my-app04-service-green.staging.svc.cluster.local:8080/hello", "http://my-app04-service-green.staging.svc.cluster.local:8080/headers"}
	fmt.Println(urlArgs)
	fmt.Printf("### Check of URL's format ### \n")
	for _, urlseq := range urlArgs {
		urlValid, err := url.ParseRequestURI(urlseq)
		if err != nil {
			fmt.Printf("error with URL: %v \n", urlseq)
			os.Exit(1)
		} else {
			fmt.Printf("valid URL: %v \n", urlValid)
		}
	}
	fmt.Printf("### URL's response starting... ### \n")
	for _, urlseq := range urlArgs {
		startime := time.Now()
		resp, err := client.Get(urlseq)
		if err != nil {
			fmt.Printf("Site: %v ERROR \n", urlseq)
			os.Exit(11)
		} else if time.Since(startime).Milliseconds() > 1500 {
			defer resp.Body.Close()
			responseTime := float64(time.Since(startime).Milliseconds())
			fmt.Printf("Site: %v -- ResponseTime: %v ms -- StatusCode: %v  take longer than expected \n", urlseq, responseTime, resp.StatusCode)
			os.Exit(62)
		} else {
			defer resp.Body.Close()
			responseTime := float64(time.Since(startime).Milliseconds())
			fmt.Printf("Site: %v -- ResponseTime: %v ms -- StatusCode: %v  \n", urlseq, responseTime, resp.StatusCode)

		}
	}

	}

