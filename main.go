package main

import (
	"encoding/json"
	"fmt"
	"github.com/esakat/observe_my_hatebu/data"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

func main() {

	// Get MyBookMark Entries from Firestore
	myEntries := GetMyHatebuEntries()

	// Wait all operation
	wg := &sync.WaitGroup{}
	for _, me := range myEntries {
		wg.Add(1)
		myEntry := me
		go func() {
			defer wg.Done()

			// Setup Request
			req, err := http.NewRequest("GET", "https://b.hatena.ne.jp/entry/jsonlite/", nil)
			if err != nil {
				log.Fatalln(err)
			}
			req.Header.Add("Accept", "application/json")
			// Do not cache
			req.Header.Add("Cache-Control", "no-store")
			q := req.URL.Query()
			q.Add("url", myEntry.URL)

			// Do not cache
			uuidObj, _ := uuid.NewUUID()
			q.Add("uuid", uuidObj.String())
			req.URL.RawQuery = q.Encode()

			// Do Request
			client := &http.Client{}
			resp, err := client.Do(req)
			defer resp.Body.Close()
			respBody, _ := ioutil.ReadAll(resp.Body)

			// Assign response to struct
			var entry data.HatebuEntry
			err = json.Unmarshal(respBody, &entry)
			if err != nil {
				log.Fatal(err)
			}

			// Filtering response
			entry.FilterNoComment().FilterOldComment(myEntry.UpdateTimestamp)

			if len(entry.Bookmarks) > 0 {
				log.Printf("%d new entries have been registered for %s.\n", len(entry.Bookmarks), entry.Title)
				// MyEntry Update
				myEntry.UpdateTimestamp = entry.Bookmarks[0].Timestamp
				UpdateMyHatebuEntry(myEntry)

				// Push Slack Notification Message
				for _, bookmark := range entry.Bookmarks {

					// create write text
					var tagMessage string
					if len(bookmark.Tags) > 0 {
						tagMessage = ":label:" + strings.Join(bookmark.Tags, ", :label:")
					}
					writeText := fmt.Sprintf("%s %s", bookmark.Comment, tagMessage)

					// create icon url
					iconUrl := fmt.Sprintf("https://cdn.profile-image.st-hatena.com/users/%s/profile.png", bookmark.User)

					// Push to Firestore
					PushSlackNotifyMessage(bookmark.User, writeText, myEntry.ThreadTimestamp, config.ChannelID, iconUrl)
				}
			} else {
				log.Printf("No new entries have been registered for %s.\n", entry.Title)

				// Delete too old my bookmark
				t, _ := time.ParseInLocation("2006/01/02 15:04", myEntry.UpdateTimestamp, data.JST)
				now := time.Now().Add(-24 * time.Hour)
				if t.Before(now) {
					DeleteMyHatebuEntry(myEntry)
					log.Printf("%s is too old, it was deleted\n", entry.Title)
				}
			}
		}()
	}
	wg.Wait()
}
