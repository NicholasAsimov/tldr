package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const baseURL = "https://raw.githubusercontent.com/tldr-pages/tldr/master/pages/common/%s.md"

const (
	ANSI_CLEAR     = "\033[0m"
	ANSI_BOLD      = "\033[1m"
	ANSI_UNDERLINE = "\033[4m"
	INDENT         = "  "
)

// TODO add caching
func getPage(name string) (io.Reader, error) {
	url := fmt.Sprintf(baseURL, name)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("can't get page: %s", err)
	}

	return resp.Body, nil
}

func render(page io.Reader) error {
	contentRaw, err := ioutil.ReadAll(page)
	if err != nil {
		return fmt.Errorf("can't read page: %s", err)
	}

	// Print initial newline
	fmt.Printf("\n")

	lines := strings.Split(string(contentRaw), "\n")

	for _, line := range lines {
		// fmt.Printf("<%q>", line)
		if line == "" {
			fmt.Printf("\n")
			continue
		}

		block := line[0]
		text := line[2:len(line)]

		fmt.Printf(INDENT)

		switch block {
		// Command name
		case '#':
			fmt.Println(ANSI_UNDERLINE + ANSI_BOLD + text + ANSI_CLEAR)
		// Command description
		case '>':
			fmt.Println(ANSI_BOLD + text + ANSI_CLEAR)
		// Example description
		case '-':
			fmt.Printf(line)
		// Example code
		// case '`':
		// 	fmt.Println(line)
		default:
			fmt.Println(line)
		}
	}

	return nil
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

	if err := render(page); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
