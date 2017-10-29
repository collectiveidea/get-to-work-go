package harvest

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
)

func mockResponse(paths ...string) *httptest.Server {
	parts := []string{".", "testdata"}
	filename := filepath.Join(append(parts, paths...)...)

	mockData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write(mockData)
	}))
}

func mockRedirectResponse(paths ...string) *httptest.Server {
	parts := []string{".", "testdata"}
	filename := filepath.Join(append(parts, paths...)...)

	mockData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" || r.Method == "PUT" {
			rw.Header().Set("Location", "/redirect/123456")
			rw.Write([]byte{})
		} else {
			rw.Write(mockData)
		}
	}))
}

func mockDynamicPathResponse() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" || r.Method == "PUT" {
			rw.Header().Set("Location", r.URL.Path+"-"+r.Method)
			rw.Write([]byte{})
			rw.Write([]byte{})
		} else {
			// Build the path for the dynamic content
			parts := []string{".", "testdata"}
			parts = append(parts, strings.Split(strings.TrimPrefix(r.URL.Path, "/"), "/")...)
			// Remove security strings
			queryStringPart := r.URL.RawQuery
			if queryStringPart != "" {
				parts[len(parts)-1] = fmt.Sprintf("%s-%x", parts[len(parts)-1], md5.Sum([]byte(queryStringPart)))
			}
			parts[len(parts)-1] = parts[len(parts)-1] + ".json"
			filename := filepath.Join(parts...)

			if _, err := os.Stat(filename); os.IsNotExist(err) {
				http.Error(rw, fmt.Sprintf("%s doesn't exist. Create it with the mock you'd like to use.\n Args were: %s", filename, queryStringPart), http.StatusNotFound)
				return
			}

			mockData, err := ioutil.ReadFile(filename)
			if err != nil {
				log.Fatal(err)
			}
			rw.Write(mockData)
		}

	}))
}

func mockErrorResponse(code int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		http.Error(rw, "An error occurred", code)
	}))
}
