package data

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestFilterNoComment(t *testing.T) {
	actual := &HatebuEntry{
		EntryURL: "https:foobar.com",
		Screenshot:   "testScreenshot",
		Bookmarks: []Bookmark{
			{
				Timestamp: "yyyymmddmmss",
				User:      "foo",
				Tags:      nil,
				Comment:   "",
			},
			{
				Timestamp: "yyyymmddmmss",
				User:      "bar",
				Tags:      nil,
				Comment:   "",
			},
			{
				Timestamp: "yyyymmddmmss",
				User:      "hoge",
				Tags:      []string{
					"test1",
					"test2",
				},
				Comment:   "exists comment",
			},
		},
		Title: "testTitle",
		URL: "testURL",
		RequestedURL: "testRequestURL",
		Eid: "testEid",
		Count: 3,
	}

	expected := &HatebuEntry{
		EntryURL: "https:foobar.com",
		Screenshot:   "testScreenshot",
		Bookmarks: []Bookmark{
			{
				Timestamp: "yyyymmddmmss",
				User:      "hoge",
				Tags:      []string{
					"test1",
					"test2",
				},
				Comment:   "exists comment",
			},
		},
		Title: "testTitle",
		URL: "testURL",
		RequestedURL: "testRequestURL",
		Eid: "testEid",
		Count: 3,
	}

	actual.FilterNoComment()

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("Hogefunc differs: (-got +want)\n%s", diff)
	}
}


func TestFilterOldComment(t *testing.T) {
	actual := &HatebuEntry{
		EntryURL: "https:foobar.com",
		Screenshot:   "testScreenshot",
		Bookmarks: []Bookmark{
			{
				Timestamp: "yyyymmddmmss",
				User:      "foo",
				Tags:      nil,
				Comment:   "",
			},
			{
				Timestamp: "yyyymmddmmss",
				User:      "bar",
				Tags:      nil,
				Comment:   "",
			},
			{
				Timestamp: "yyyymmddmmss",
				User:      "hoge",
				Tags:      []string{
					"test1",
					"test2",
				},
				Comment:   "exists comment",
			},
		},
		Title: "testTitle",
		URL: "testURL",
		RequestedURL: "testRequestURL",
		Eid: "testEid",
		Count: 3,
	}

	expected := &HatebuEntry{
		EntryURL: "https:foobar.com",
		Screenshot:   "testScreenshot",
		Bookmarks: []Bookmark{
			{
				Timestamp: "yyyymmddmmss",
				User:      "hoge",
				Tags:      []string{
					"test1",
					"test2",
				},
				Comment:   "exists comment",
			},
		},
		Title: "testTitle",
		URL: "testURL",
		RequestedURL: "testRequestURL",
		Eid: "testEid",
		Count: 3,
	}

	actual.FilterOldComment("")

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("Hogefunc differs: (-got +want)\n%s", diff)
	}
}
