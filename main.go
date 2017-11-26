package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const baseURL = "https://raw.githubusercontent.com/tldr-pages/tldr/master/pages/common/%s.md"

func getPage(name string) (io.Reader, error) {
	url := fmt.Sprintf(baseURL, name)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("can't get page: %s", err)
	}

	return resp.Body, nil
}

func main() {
	var pageName string

	if len(os.Args) > 1 {
		pageName = os.Args[1]
	}

	if pageName == "" {
		fmt.Println("Specify page name")
		os.Exit(1)
	}

	page, err := getPage(pageName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = io.Copy(os.Stdout, page)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
