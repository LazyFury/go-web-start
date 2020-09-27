package api

import (
	"encoding/xml"
	"net/http"

	"github.com/labstack/echo/v4"
)

type rss struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Itunes  string   `xml:"xmlns:itunes,attr"`
	Content string   `xml:"xmlns:content,attr"`

	Channel podcast `xml:"channel"`
}

type podcast struct {
	XMLName xml.Name `xml:"channel"`

	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Language    string `xml:"language"`
	CopyRight   string `xml:"copyRight"`
	Description string `xml:"description"`
	itunes

	Item []item `xml:"item"`
}
type itunes struct {
	Author   string   `xml:"itunes:author"`
	Type     string   `xml:"itunes:type"`
	Owner    owner    `xml:"itunes:owner"`
	Image    string   `xml:"itunes:image"`
	Category category `xml:"itunes:category"`
	Explicit bool     `xml:"itunes:explicit"`
}

type category struct {
	Text string `xml:"text,attr"`
}

type owner struct {
	Name  string `xml:"itunes:name"`
	Email string `xml:"itunes:email"`
}

type item struct {
	Title       string    `xml:"title"`
	Description cdata     `xml:"description"`
	Link        string    `xml:"link,omitempty"`
	Enclosure   enclosure `xml:"enclosure"`
	GUID        string    `xml:"guid"`
	PubDate     string    `xml:"pubDate"`
	itemItunes
}

type itemItunes struct {
	EpisodeType string `xml:"itunes:episodeType"` //full,trailer预告
	Episode     string `xml:"itunes:episode"`
	Season      string `xml:"itunes:season"`
	Duration    int    `xml:"itunes:duration"`
	Explicit    bool   `xml:"itunes:explicit"`
}

type enclosure struct {
	Length int    `xml:"length,attr"`
	Type   string `xml:"type,attr"`
	URL    string `xml:"url,attr"`
}

type cdata struct {
	Content string `xml:",cdata"`
}

func newRss(p podcast) *rss {
	return &rss{
		Itunes:  "http://www.itunes.com/dtds/podcast-1.0.dtd",
		Content: "http://purl.org/rss/1.0/modules/content/",
		Version: "2.0",
		Channel: p,
	}
}

func podcastRouter(g *echo.Group) {
	pod := g.Group("/podcast")

	pod.GET("/test", testPodcast)
}

func testPodcast(c echo.Context) error {

	pod := newRss(podcast{
		Title:       "test",
		Description: "desc",
		Link:        "http://go.abadboy.cn",
		Language:    "zh_CN",
		CopyRight:   "suke",
		itunes: itunes{
			Author: "suke",
			Type:   "serial",
			Owner: owner{
				Name:  "2568597007@qq.com",
				Email: "suke@qq.com",
			},
			Image:    "http://wx2.sinaimg.cn/mw600/0085KTY1gy1ghkic07i5yj30qo0qo41n.jpg",
			Explicit: false,
			Category: category{"测试"},
		},
		Item: []item{
			{
				Title: "测试项目标题",
				Enclosure: enclosure{
					Length: 1000,
					Type:   "audio/mpeg",
					URL:    "http://go.abadboy.cn/4.mp4",
				},
				GUID:    "2344567890",
				PubDate: "Tue, 07 May 2019 12:00:00 GMT",
				itemItunes: itemItunes{
					EpisodeType: "full",
					Episode:     "4",
					Season:      "2",
					Duration:    3456789,
					Explicit:    false,
				},
				Description: cdata{"hello"}},
		},
	})

	return c.XMLPretty(http.StatusOK, pod, "	")
}
