package counter

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

const goWordRegExp = `\b(Go)\b`

// isValidUrl tests a string to determine if it is a well-structured url or not
func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

// getReader detects source type and returns valid reader
func getReader(s string) io.ReadCloser {

	// Create HTTP reader
	if isValidUrl(s) {
		resp, err := http.Get(s)
		if err == nil {
			return resp.Body
		}
	}

	// Create file reader
	f, err := os.Open(s)
	if err == nil {
		return f
	}

	// Create string reader for test purposes
	return ioutil.NopCloser(strings.NewReader(s))
}

// CountGoWords returns number of words "Go" in a text source
func CountGoWords(source string) int {

	// Initialize source reader
	reader := getReader(source)
	text, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Println(err)
		return 0
	}
	defer reader.Close()

	// Find substrings with regular expressions
	re := regexp.MustCompile(goWordRegExp)
	loc := re.FindAll(text, -1)
	return len(loc)
}
