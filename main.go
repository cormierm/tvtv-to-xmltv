package main

import (
	"log"
	"net/http"
	"text/template"
	"tvtvToXmltv/tvtv"
	"tvtvToXmltv/xmltv"
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

	tvtvListing, err := tvtv.FetchListing()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	xml := xmltv.TvtvToXMLTV(tvtvListing)

	err = tpl.ExecuteTemplate(w, "xmltv.goxml", xml)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
