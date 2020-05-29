package xmltv

import "encoding/xml"

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
