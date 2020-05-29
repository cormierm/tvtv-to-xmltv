package main

import (
	"encoding/json"
	"io/ioutil"
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

	tvtvList := getTvtvListing()

	xml := xmltv.TvtvToXMLTV(tvtvList)

	err := tpl.ExecuteTemplate(w, "xmltv.goxml", xml)
	if err != nil {
		log.Println(err)
	}
}

func getTvtvListing() tvtv.Tvtv {
	tvtvUrl := "https://tvtv.ca/tvm/t/tv/v4/lineups/3003/listings/grid?start=2020-05-28&end=2020-05-30"

	resp, err := http.Get(tvtvUrl)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var tvtvList tvtv.Tvtv
	err = json.Unmarshal(body, &tvtvList)
	if err != nil {
		log.Println(err)
	}

	return tvtvList
}
