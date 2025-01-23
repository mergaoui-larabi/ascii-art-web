package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	asciigenerator "ascii-art-web/ascii-generator"
)

// func getHandler(w http.ResponseWriter, r *http.Request) {
// 	data := r.URL.Query().Get("data")
// 	fmt.Fprintf(w, "<h1>%s</h1>", data)
// }

func postHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "howa error")
	}

	result := asciigenerator.AsciiGenerator(r.FormValue("ascii-input"), "standard")

	fmt.Println(result)
	spl := strings.Split(result, "\r\n")
	for i := 0; i < len(spl); i++ {
		fmt.Fprintln(w, spl[i])
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	// serving the first html file
	http.HandleFunc("/", rootHandler)

	// http.HandleFunc("/get", getHandler)
	http.HandleFunc("/ascii-art", postHandler)

	log.Println(http.ListenAndServe(":8080", nil))
}
