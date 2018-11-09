package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type WCService struct {
}

func (wc *WCService) handleWordCount(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	s := q.Get("s")

	s = strings.TrimSpace(s)

	res := 0

	log.Printf("[word count] received %s", s)
	if s != "" {
		res = res + 1
		for _, c := range s {
			if c == 32 {
				res = res + 1
			}
		}
	}
	io.WriteString(w, strconv.Itoa(res))
}

func main() {
	mux := http.NewServeMux()

	wc := WCService{}

	mux.HandleFunc("/word", wc.handleWordCount)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
