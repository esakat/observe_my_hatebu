package main

import (
	"github.com/kelseyhightower/envconfig"
	"log"
	"time"
)

type Config struct {
	RedisHatebuEntryAddr  string        `required:"true" split_words:"true"`
	RedisSlackMessageAddr string        `required:"true" split_words:"true"`
	RedisHatebuEntryDB    int           `required:"true" default:"0" split_words:"true"`
	RedisSlackMessageDB   int           `required:"true" default:"1" split_words:"true"`
	PushCommentKey        string        `default:"notify-queue" split_words:"true"`
	UpdateDuration        time.Duration `default:"3m" split_words:"true"`
	BotToken              string        `required:"true" split_words:"true"`
	ChannelID             string        `required:"true" split_words:"true"`
}

var config Config

func init() {
	err := envconfig.Process("app", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	createRedisClient()
}
