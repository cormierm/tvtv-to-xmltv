package tvtv

var DateTimeLayout = "2006-01-02 15:04:05"

type Tvtv []struct {
	Channel  Channel   `json:"channel"`
	Listings []Listing `json:"listings"`
}

type Channel struct {
		Name             string `json:"name"`
		Number           string `json:"number"`
		ChannelNumber    int    `json:"channelNumber"`
		SubChannelNumber int    `json:"subChannelNumber"`
		StationID        int    `json:"stationID"`
		Callsign         string `json:"callsign"`
		Network          string `json:"network"`
		StationType      string `json:"stationType"`
		NTSCTSID         int    `json:"NTSC_TSID"`
		DTVTSID          int    `json:"DTV_TSID"`
		Twitter          string `json:"Twitter"`
		WebLink          string `json:"webLink"`
		LogoFilename     string `json:"logoFilename"`
		StationHD        bool   `json:"stationHD"`
}

type Listing struct {
		ListingID        int    `json:"listingID"`
		ListDateTime     string `json:"listDateTime"`
		Duration         int    `json:"duration"`
		ShowID           int    `json:"showID"`
		SeriesID         int    `json:"seriesID"`
		ShowName         string `json:"showName"`
		EpisodeTitle     string `json:"episodeTitle"`
		EpisodeNumber    string `json:"episodeNumber"`
		Parts            int    `json:"parts"`
		PartNum          int    `json:"partNum"`
		SeriesPremiere   bool   `json:"seriesPremiere"`
		SeasonPremiere   bool   `json:"seasonPremiere"`
		SeriesFinale     bool   `json:"seriesFinale"`
		SeasonFinale     bool   `json:"seasonFinale"`
		Repeat           bool   `json:"repeat"`
		New              bool   `json:"new"`
		Rating           string `json:"rating"`
		Captioned        bool   `json:"captioned"`
		Educational      bool   `json:"educational"`
		BlackWhite       bool   `json:"blackWhite"`
		Subtitled        bool   `json:"subtitled"`
		Live             bool   `json:"live"`
		Hd               bool   `json:"hd"`
		DescriptiveVideo bool   `json:"descriptiveVideo"`
		InProgress       bool   `json:"inProgress"`
		ShowTypeID       string `json:"showTypeID"`
		BreakoutLevel    int    `json:"breakoutLevel"`
		ShowType         string `json:"showType"`
		Year             string `json:"year"`
		Guest            string `json:"guest"`
		Cast             string `json:"cast"`
		Director         string `json:"director"`
		StarRating       int    `json:"starRating"`
		Description      string `json:"description"`
		League           string `json:"league"`
		Team1ID          int    `json:"team1ID"`
		Team2ID          int    `json:"team2ID"`
		Team1            string `json:"team1"`
		Team2            string `json:"team2"`
		Event            string `json:"event"`
		Location         string `json:"location"`
		ShowPicture      string `json:"showPicture"`
		Artwork          struct {
			Showcard    interface{} `json:"showcard"`
			Titlecard   interface{} `json:"titlecard"`
			Poster      interface{} `json:"poster"`
			Episodic    interface{} `json:"episodic"`
			MovieStill  interface{} `json:"movieStill"`
			MoviePoster interface{} `json:"moviePoster"`
		} `json:"artwork"`
		ShowHost string `json:"showHost"`
}
