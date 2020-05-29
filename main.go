package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"tvtvToXmltv/tvtv"
)

func getTvtvListingToXmlTV(w http.ResponseWriter, req *http.Request) {
	log.Printf("[%v] Requesting TvtvListToXmlTV\n", req.RemoteAddr)
	xml := getXml()
	fmt.Fprint(w, xml)
}

func main() {
	http.HandleFunc("/", getTvtvListingToXmlTV)
	http.ListenAndServe(":8080", nil)
}

func getXml() string {
	tvtvList := getTvtvListing()

	type XmlChannel struct {
		XMLName xml.Name `xml:"channel"`
		Id      string   `xml:"id,attr"`
		Name    []string   `xml:"display-name"`
		Icon        struct {
			Text string `xml:",chardata"`
			Src  string `xml:"src,attr"`
		} `xml:"icon"`
	}

	type XmlProgramme struct {
		XMLName xml.Name `xml:"programme"`
		Start   string   `xml:"start,attr"`
		Stop    string   `xml:"stop,attr"`
		Channel string   `xml:"channel,attr"`
		Title   struct {
			Text string `xml:",chardata"`
			Lang string `xml:"lang,attr"`
		} `xml:"title"`
		SubTitle struct {
			Text string `xml:",chardata"`
			Lang string `xml:"lang,attr"`
		} `xml:"sub-title"`
		Desc struct {
			Text string `xml:",chardata"`
			Lang string `xml:"lang,attr"`
		} `xml:"desc"`
		Date     string `xml:"date"`
		Category []struct {
			Text string `xml:",chardata"`
			Lang string `xml:"lang,attr"`
		} `xml:"category"`
		EpisodeNum []struct {
			Text   string `xml:",chardata"`
			System string `xml:"system,attr"`
		} `xml:"episode-num"`
		Audio struct {
			Text   string `xml:",chardata"`
			Stereo string `xml:"stereo"`
		} `xml:"audio"`
		PreviouslyShown struct {
			Text  string `xml:",chardata"`
			Start string `xml:"start,attr"`
		} `xml:"previously-shown"`
		Subtitles struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"subtitles"`
	}

	type XmlTv struct {
		XMLName           xml.Name `xml:"tv"`
		SourceInfoURL     string   `xml:"source-info-url,attr"`
		SourceInfoName    string   `xml:"source-info-name,attr"`
		GeneratorInfoName string   `xml:"generator-info-name,attr"`
		GeneratorInfoURL  string   `xml:"generator-info-url,attr"`
		Channels []*XmlChannel `xml:"channel"`
		Programme []*XmlProgramme `xml:"programme"`
	}

	tv := &XmlTv{
		SourceInfoURL: "http://localhost",
		SourceInfoName: "My Tvtv to XmlTV",
	}

	for _, channel := range tvtvList {
		id := "my.xml." + channel.Channel.Number
		names := []string{
			channel.Channel.Name,
			channel.Channel.Number,
			channel.Channel.Callsign,
		}

		xmlChannel := &XmlChannel{
			XMLName: xml.Name{},
			Id:      id,
			Name:    names,
			Icon: struct {
				Text string `xml:",chardata"`
				Src  string `xml:"src,attr"`
			}{
				Src: fmt.Sprintf("https://cdn.tvpassport.com/image/station/100x100/%v",channel.Channel.LogoFilename),
			},
		}

		for _, program := range channel.Listings {
			xmlProgramme := &XmlProgramme{
				XMLName: xml.Name{},
				Start:   convertTimestamp(program.ListDateTime),
				Stop:    xmlStopTimestamp(program.ListDateTime, program.Duration),
				Channel: id,
				Title: struct {
					Text string `xml:",chardata"`
					Lang string `xml:"lang,attr"`
				}{
					sanitizeString(program.ShowName),
				"en",
				},
				SubTitle: struct {
					Text string `xml:",chardata"`
					Lang string `xml:"lang,attr"`
				}{Text: sanitizeString(program.EpisodeTitle), Lang: "en"},
				Desc: struct {
					Text string `xml:",chardata"`
					Lang string `xml:"lang,attr"`
				}{
					sanitizeString(program.Description) + ".",
					"en",
				},
				Date:       "2000-01-01",
				Category:   nil,
				EpisodeNum: nil,
				Audio: struct {
					Text   string `xml:",chardata"`
					Stereo string `xml:"stereo"`
				}{
					Stereo: "stereo",
				},
				PreviouslyShown: struct {
					Text  string `xml:",chardata"`
					Start string `xml:"start,attr"`
				}{},
				Subtitles: struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				}{
					Type: "teletext",
				},
			}
			tv.Programme = append(tv.Programme, xmlProgramme)
		}

		tv.Channels = append(tv.Channels, xmlChannel)
	}

	out, _ := xml.MarshalIndent(&tv, "", "    ")

	return 	"<?xml version=\"1.0\" encoding=\"ISO-8859-1\"?><!DOCTYPE tv SYSTEM \"xmltv.dtd\">" + string(out)
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

func convertTimestamp(timestamp string) string {
	layoutTvtv := "2006-01-02 15:04:05"
	layoutXml := "20060102150405 -0700"
	t, err := time.Parse(layoutTvtv, timestamp)
	if err != nil {
		log.Println("xmlDatetime error:", err)
	}

	return t.Format(layoutXml)
}

func xmlStopTimestamp(timestamp string, minsToAdd int) string {
	layoutTvtv := "2006-01-02 15:04:05"
	layoutXml := "20060102150405 -0700"
	t, err := time.Parse(layoutTvtv, timestamp)
	if err != nil {
		log.Println("xmlDatetime error:", err)
	}
	t = t.Add(time.Duration(minsToAdd) * time.Minute)

	return t.Format(layoutXml)
}

func sanitizeString(text string) string {
	return text
}