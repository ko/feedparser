package feedparser

import (
	"encoding/json"
	"fmt"
)

type ItunesSearchResultsItem struct {
	ArtistID               int      `json:"artistId"`
	ArtistName             string   `json:"artistName"`
	ArtistViewURL          string   `json:"artistViewUrl"`
	ArtworkURL100          string   `json:"artworkUrl100"`
	ArtworkURL30           string   `json:"artworkUrl30"`
	ArtworkURL60           string   `json:"artworkUrl60"`
	ArtworkURL600          string   `json:"artworkUrl600"`
	CollectionCensoredName string   `json:"collectionCensoredName"`
	CollectionExplicitness string   `json:"collectionExplicitness"`
	CollectionHdPrice      int      `json:"collectionHdPrice"`
	CollectionID           int      `json:"collectionId"`
	CollectionName         string   `json:"collectionName"`
	CollectionPrice        int      `json:"collectionPrice"`
	CollectionViewURL      string   `json:"collectionViewUrl"`
	ContentAdvisoryRating  string   `json:"contentAdvisoryRating"`
	Country                string   `json:"country"`
	Currency               string   `json:"currency"`
	FeedURL                string   `json:"feedUrl"`
	GenreIds               []string `json:"genreIds"`
	Genres                 []string `json:"genres"`
	Kind                   string   `json:"kind"`
	PrimaryGenreName       string   `json:"primaryGenreName"`
	ReleaseDate            string   `json:"releaseDate"`
	TrackCensoredName      string   `json:"trackCensoredName"`
	TrackCount             int      `json:"trackCount"`
	TrackExplicitness      string   `json:"trackExplicitness"`
	TrackHdPrice           int      `json:"trackHdPrice"`
	TrackHdRentalPrice     int      `json:"trackHdRentalPrice"`
	TrackID                int      `json:"trackId"`
	TrackName              string   `json:"trackName"`
	TrackPrice             int      `json:"trackPrice"`
	TrackRentalPrice       int      `json:"trackRentalPrice"`
	TrackViewURL           string   `json:"trackViewUrl"`
	WrapperType            string   `json:"wrapperType"`
}

type ItunesSearchResults struct {
	ResultCount int                       `json:"resultCount"`
	Results     []ItunesSearchResultsItem `json:"results"`
}

func SearchUrlGenerator(query string) string {
	url := fmt.Sprintf("https://itunes.apple.com/search?term=%s&entity=podcast", query)
	return url
}

func Search(query string) []ItunesSearchResultsItem {
	url := SearchUrlGenerator(query)
	contents := GetFeed(url)
	results := FeedToItunesSearchResults(contents)
	podcasts := FilterOnlyPodcasts(results)

	return podcasts.Results
}

func SearchResultsItemsToJson(results []ItunesSearchResultsItem) ([]byte, error) {
	barray, err := json.Marshal(results)
	return barray, err
}

func FeedToItunesSearchResults(contents string) ItunesSearchResults {
	var results ItunesSearchResults

	json.Unmarshal([]byte(contents), &results)

	return results
}

func FilterOnlyPodcasts(results ItunesSearchResults) ItunesSearchResults {
	var podcasts ItunesSearchResults
	var podcastCount int
	var podcastList []ItunesSearchResultsItem

	podcastCount = 0

	for index := range results.Results {
		podcastCount = podcastCount + 1
		if results.Results[index].Kind == "podcast" {
			podcastList = append(podcastList, results.Results[index])
		}
	}

	podcasts.Results = podcastList
	podcasts.ResultCount = podcastCount

	return podcasts
}
