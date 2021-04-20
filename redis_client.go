package main

import (
	"context"
	"encoding/json"
	"github.com/esakat/observe_my_hatebu/data"
	"github.com/go-redis/redis/v8"
	"log"
)

var ctx = context.Background()

var rdbHatebuEntry *redis.Client
var rdbSlackMessage *redis.Client

func createRedisClient() {
	rdbHatebuEntry = redis.NewClient(&redis.Options{
		Addr:     config.RedisHatebuEntryAddr,
		Password: "",
		DB:       config.RedisHatebuEntryDB,
	})

	rdbSlackMessage = redis.NewClient(&redis.Options{
		Addr:     config.RedisSlackMessageAddr,
		Password: "",
		DB:       config.RedisSlackMessageDB,
	})
}

func GetMyHatebuEntries() []*data.MyEntry {

	var entries []*data.MyEntry

	keys, err := rdbHatebuEntry.Keys(ctx, "*").Result()
	if err == redis.Nil {
		log.Println("key does not exist")
		return entries
	} else if err != nil {
		panic(err)
	}

	for _, key := range keys {
		result, err := rdbHatebuEntry.HGetAll(ctx, key).Result()
		if err != nil {
			panic(err)
		}

		entry := &data.MyEntry{
			EntryID:         key,
			URL:             result["url"],
			ThreadTimestamp: result["thread_timestamp"],
			UpdateTimestamp: result["update_timestamp"],
		}
		entries = append(entries, entry)
	}
	return entries
}

func UpdateMyHatebuEntry(entry *data.MyEntry) {
	var m = make(map[string]interface{})
	m["update_timestamp"] = entry.UpdateTimestamp
	_, err := rdbHatebuEntry.HMSet(ctx, entry.EntryID, m).Result()
	if err != nil {
		panic(err)
	}
}

func DeleteMyHatebuEntry(entry *data.MyEntry) {
	_, err := rdbHatebuEntry.Del(ctx, entry.EntryID).Result()
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
	slackMessage, _ := json.Marshal(sm)

	_, err := rdbSlackMessage.LPush(ctx, config.PushCommentKey, slackMessage).Result()
	if err != nil {
		panic(err)
	}
}
