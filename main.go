package main

import (
	"fmt"
	"github.com/cormierm/TvtvToXmlTV/tvtv"
	"github.com/cormierm/TvtvToXmlTV/xmltv"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.goxml"))
}

func main() {
	http.HandleFunc("/", xmltvHandlerFunc)
	http.ListenAndServe(":8080", nil)
}

func xmltvHandlerFunc(w http.ResponseWriter, req *http.Request) {
	log.Printf("[%v] Requesting TvtvListToXmlTV\n", req.RemoteAddr)

	query := req.URL.Query()
	days := query.Get("days")
	fmt.Println(days)
	if days == "" {
		http.Error(w, "Required parameter days is missing.", http.StatusUnprocessableEntity)
		return
	}
	intDays, err := strconv.Atoi(days)
	if err != nil {
		http.Error(w, "Invalid days parameter. Must be integer.", http.StatusUnprocessableEntity)
		return
	}

	location := query.Get("location")
	if location == "" {
		http.Error(w, "Required parameter location is missing.", http.StatusUnprocessableEntity)
		return
	}

	tvtvListing, err := tvtv.FetchListing(location, intDays)
	if err != nil {
		http.Error(w, "Error tvtv.Fetching: " + err.Error(), http.StatusInternalServerError)
		return
	}

	xml := xmltv.TvtvToXMLTV(tvtvListing)

	err = tpl.ExecuteTemplate(w, "xmltv.goxml", xml)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error ExecuteTemplate: " + err.Error(), http.StatusInternalServerError)
		return
	}
}
