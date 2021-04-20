package data

import (
	"time"
)

var JST, _ = time.LoadLocation("Asia/Tokyo")

type Bookmark struct {
	Timestamp string   `json:"timestamp"`
	User      string   `json:"user"`
	Tags      []string `json:"tags"`
	Comment   string   `json:"comment"`
}

type HatebuEntry struct {
	EntryURL     string     `json:"entry_url"`
	Screenshot   string     `json:"screenshot"`
	Bookmarks    []Bookmark `json:"bookmarks"`
	Title        string     `json:"title"`
	URL          string     `json:"url"`
	RequestedURL string     `json:"requested_url"`
	Eid          string     `json:"eid"`
	Count        int        `json:"count"`
}

func (this *HatebuEntry) FilterNoComment() *HatebuEntry {
	var tmp []Bookmark
	for _, v := range this.Bookmarks {
		if len(v.Comment) > 0 {
			tmp = append(tmp, v)
		}
	}
	this.Bookmarks = tmp
	return this
}

func (this *HatebuEntry) FilterOldComment(beforeExecutedTimeString string) *HatebuEntry {
	beforeExecutedTime, _ := time.ParseInLocation("2006/01/02 15:04", beforeExecutedTimeString, JST)
	var tmp []Bookmark
	for _, v := range this.Bookmarks {
		t, _ := time.ParseInLocation("2006/01/02 15:04", v.Timestamp, JST)
		if t.After(beforeExecutedTime) {
			tmp = append(tmp, v)
		}
	}
	this.Bookmarks = tmp
	return this
}
