package tvtv

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func FetchListing() (Tvtv, error) {
	tvtvUrl := generateTvtvUrl("3003", 1)

	log.Printf("Fetching tvtv listing: %s\n", tvtvUrl)
	resp, err := http.Get(tvtvUrl)
	if err != nil {
		log.Printf("Error fetching url: %s %s", tvtvUrl, err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error parsing body:", err)
		return nil, err
	}

	var listing Tvtv

	err = json.Unmarshal(body, &listing)
	if err != nil {
		log.Println("Error unmarshalling tvtv response:", err)
		return nil, err
	}

	return listing, nil
}

func generateTvtvUrl(locationId string, days int) string {
	urlTimeLayout := "2006-01-02 15:00:00"

	start := time.Now().Format(urlTimeLayout)
	end := time.Now().Add(time.Duration(days) * time.Hour * 24).Format(urlTimeLayout)

	params := url.Values{}
	params.Add("start", start)
	params.Add("end", end)

	tvtvUrl := &url.URL{
		Scheme:   "https",
		Host:     "tvtv.ca",
		Path:     fmt.Sprintf("/tvm/t/tv/v4/lineups/%s/listings/grid", locationId),
		RawQuery: params.Encode(),
	}

	// https://tvtv.ca/tvm/t/tv/v4/lineups/3003/listings/grid?end=2020-05-30+11%3A00%3A00&start=2020-05-29+11%3A00%3A00
	return tvtvUrl.String()
}