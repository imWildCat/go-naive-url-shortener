package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var urlMap = make(map[string]string)

func loadConfig(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), " ", 2)
		if len(parts) == 2 {
			urlMap[parts[0]] = parts[1]
		}
	}

	return scanner.Err()
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if url, ok := urlMap[path]; ok {
		http.Redirect(w, r, url, http.StatusFound)
		return
	}
	http.NotFound(w, r)
}

func main() {
	if err := loadConfig("/app/config.txt"); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", redirectHandler)

	fmt.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
