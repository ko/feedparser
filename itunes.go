package feedparser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	xj "github.com/basgys/goxml2json"
)

type Item struct {
	Author      []string `json:"author"`
	Category    string   `json:"category"`
	Description string   `json:"description"`
	Duration    string   `json:"duration"`
	Enclosure   struct {
		Length string `json:"-length"`
		Type   string `json:"-type"`
		URL    string `json:"-url"`
	} `json:"enclosure"`
	Encoded  string `json:"encoded"`
	Explicit string `json:"explicit"`
	GUID     string `json:"guid"`
	Image    struct {
		Href string `json:"-href"`
	} `json:"image"`
	Keywords string `json:"keywords"`
	Link     string `json:"link"`
	PubDate  string `json:"pubDate"`
	Subtitle string `json:"subtitle"`
	Summary  string `json:"summary"`
	Title    string `json:"title"`
}

type ItunesFeed struct {
	Rss struct {
		Atom    string `json:"-atom"`
		Content string `json:"-content"`
		Itunes  string `json:"-itunes"`
		Sy      string `json:"-sy"`
		Version string `json:"-version"`
		Channel struct {
			Author   string `json:"author"`
			Category []struct {
				Text     string `json:"-text"`
				Category struct {
					Text string `json:"-text"`
				} `json:"category"`
			} `json:"category"`
			Copyright   string `json:"copyright"`
			Description string `json:"description"`
			Explicit    string `json:"explicit"`
			Image       []struct {
				Href  string `json:"-href"`
				Link  string `json:"link"`
				Title string `json:"title"`
				URL   string `json:"url"`
			} `json:"image"`
			Items         []Item        `json:"item"`
			Keywords      string        `json:"keywords"`
			Language      string        `json:"language"`
			LastBuildDate string        `json:"lastBuildDate"`
			Link          []interface{} `json:"link"`
			New_feed_url  string        `json:"new-feed-url"`
			Owner         struct {
				Email string `json:"email"`
				Name  string `json:"name"`
			} `json:"owner"`
			PubDate         string `json:"pubDate"`
			Subtitle        string `json:"subtitle"`
			Summary         string `json:"summary"`
			Title           string `json:"title"`
			UpdateFrequency string `json:"updateFrequency"`
			UpdatePeriod    string `json:"updatePeriod"`
		} `json:"channel"`
	} `json:"rss"`
}

func GetFeed(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	byteArray, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	contents := string(byteArray[:])

	return contents
}

func GetFeedLocal() string {
	b, err := ioutil.ReadFile("test.xml")
	if err != nil {
		log.Fatal(err)
	}

	return string(b)
}

func XmlToJson(contents string) []byte {
	xml := strings.NewReader(contents)
	json, err := xj.Convert(xml)
	if err != nil {
		log.Fatal(err)
	}

	return json.Bytes()
}

func JsonToItunesFeed(byteArray []byte) ItunesFeed {
	var feed ItunesFeed
	json.Unmarshal(byteArray, &feed)

	return feed
}

func GetChannelTitle(feed ItunesFeed) string {
	return feed.Rss.Channel.Title
}

func GetChannelSubtitle(feed ItunesFeed) string {
	return feed.Rss.Channel.Subtitle
}

func GetChannelAuthor(feed ItunesFeed) string {
	return feed.Rss.Channel.Author
}

func GetChannelDescription(feed ItunesFeed) string {
	return feed.Rss.Channel.Description
}

func GetChannelImageUrl(feed ItunesFeed) string {
	return feed.Rss.Channel.Image[0].Href
}

func GetEpisodeTitles(feed ItunesFeed) []string {
	var titles []string
	var title string

	for index, _ := range feed.Rss.Channel.Items {
		title = feed.Rss.Channel.Items[index].Title
		titles = append(titles, title)
	}

	return titles
}

func GetEpisodes(feed ItunesFeed) []Item {
	return feed.Rss.Channel.Items
}

func GetEpisode(items []Item, number int) Item {
	return items[number]
}

func GetEpisodeNotes(item Item) string {
	return item.Encoded
}

func GetEpisodeImageUrl(item Item) string {
	return item.Image.Href
}
