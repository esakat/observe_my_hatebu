package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/esakat/observe_my_hatebu/data"
	"google.golang.org/api/iterator"
	"log"
)

var ctx = context.Background()
var firestoreClient *firestore.Client

func createFirestoreClient() {
	var err error
	firestoreClient, err = firestore.NewClient(ctx, config.ProjectID)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
}

func GetMyHatebuEntries() []*data.MyEntry {

	var entries []*data.MyEntry

	iter := firestoreClient.Collection(config.EntryCollectionName).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("failed to iterate: %v", err)
		}

		var entry data.MyEntry
		doc.DataTo(&entry)
		entries = append(entries, &entry)
	}

	return entries
}

func UpdateMyHatebuEntry(entry *data.MyEntry) {
	_, err := firestoreClient.Collection(config.EntryCollectionName).Doc(entry.EntryID).Set(ctx, entry)
	if err != nil {
		panic(err)
	}
}

func DeleteMyHatebuEntry(entry *data.MyEntry) {
	_, err := firestoreClient.Collection(config.EntryCollectionName).Doc(entry.EntryID).Delete(ctx)
	if err != nil {
		panic(err)
	}
}

func PushSlackNotifyMessage(username, text, threadTimestamp, channelId, iconUrl string) {
	sm := data.SlackMessage{
		UserName:        username,
		Text:            text,
		ThreadTimestamp: threadTimestamp,
		IconURL:         iconUrl,
		ChannelID:       channelId,
	}
	_, _, err := firestoreClient.Collection(config.SlackMessageCollectionName).Add(ctx, sm)
	if err != nil {
		panic(err)
	}
}
