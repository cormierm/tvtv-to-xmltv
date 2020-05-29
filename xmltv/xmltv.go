package xmltv

import (
	"encoding/xml"
	"fmt"
	"log"
	"time"
	"tvtvToXmltv/tvtv"
)

var DateTimeLayout = "20060102150405 -0700"

type Channel struct {
	XMLName xml.Name `xml:"channel"`
	Id      string   `xml:"id,attr"`
	Name    []string `xml:"display-name"`
	Icon    struct {
		Text string `xml:",chardata"`
		Src  string `xml:"src,attr"`
	} `xml:"icon"`
}

type Programme struct {
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

type Tv struct {
	XMLName           xml.Name     `xml:"tv"`
	SourceInfoURL     string       `xml:"source-info-url,attr"`
	SourceInfoName    string       `xml:"source-info-name,attr"`
	GeneratorInfoName string       `xml:"generator-info-name,attr"`
	GeneratorInfoURL  string       `xml:"generator-info-url,attr"`
	Channels          []*Channel   `xml:"channel"`
	Programme         []*Programme `xml:"programme"`
}

func TvtvToXMLTV(tvtvList tvtv.Tvtv) string {
	tv := &Tv{
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

		xmlChannel := &Channel{
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
			xmlProgramme := &Programme{
				XMLName: xml.Name{},
				Start:   convertToXmlTvTimestamp(program.ListDateTime, tvtv.DateTimeLayout),
				Stop:    xmlStopTimestamp(program.ListDateTime, program.Duration),
				Channel: id,
				Title: struct {
					Text string `xml:",chardata"`
					Lang string `xml:"lang,attr"`
				}{
					program.ShowName,
					"en",
				},
				SubTitle: struct {
					Text string `xml:",chardata"`
					Lang string `xml:"lang,attr"`
				}{Text: program.EpisodeTitle, Lang: "en"},
				Desc: struct {
					Text string `xml:",chardata"`
					Lang string `xml:"lang,attr"`
				}{
					program.Description + ".",
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

	xmltvXml, _ := xml.MarshalIndent(&tv, "", "    ")

	return string(xmltvXml)
}


func convertToXmlTvTimestamp(timestamp string, timeLayout string) string {
	t, err := time.Parse(timeLayout, timestamp)
	if err != nil {
		log.Println("xmlDatetime error:", err)
	}

	return t.Format(DateTimeLayout)
}

func xmlStopTimestamp(timestamp string, minsToAdd int) string {
	t, err := time.Parse(tvtv.DateTimeLayout, timestamp)
	if err != nil {
		log.Println("xmlDatetime error:", err)
	}
	t = t.Add(time.Duration(minsToAdd) * time.Minute)

	return t.Format(DateTimeLayout)
}

