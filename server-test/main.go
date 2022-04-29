package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
)

func main() {
	http.HandleFunc("/", topTenMostUsedWord)

	log.Println("/ POST REQUEST")
	log.Println("starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

func topTenMostUsedWord(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		http.Error(w, "Method not allow", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "text/plain" {
		http.Error(w, "Content-Type not allow", http.StatusMethodNotAllowed)
		return
	}

	requestData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	requestString := strings.ToLower(string(requestData))
	requestString = strings.ReplaceAll(requestString, ".", "")
	requestString = strings.ReplaceAll(requestString, ",", "")

	wordlist := strings.Fields(requestString)

	m := make(map[string]int)
	for _, word := range wordlist {
		m[word]++
	}

	type kv struct {
		Key   string
		Value int
	}

	ss := make([]kv, 0)
	for key, value := range m {
		ss = append(ss, kv{key, value})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	mostused := 10
	if len(ss) < 10 {
		mostused = len(ss)
	}

	response := "The top ten most-used words are\n"
	for i := 0; i < mostused; i++ {
		response += fmt.Sprintf("%d. %s: %d\n", i+1, ss[i].Key, ss[i].Value)
	}

	fmt.Fprintln(w, response)
}
