package main

import (
	"github.com/kelseyhightower/envconfig"
	"log"
	"time"
)

type Config struct {
	UpdateDuration             time.Duration `default:"3m" split_words:"true"`
	BotToken                   string        `required:"true" split_words:"true"`
	ChannelID                  string        `required:"true" split_words:"true"`
	EntryCollectionName        string        `required:"true" split_words:"true"`
	SlackMessageCollectionName string        `required:"true" split_words:"true"`
	ProjectID string `required:"true" split_words:"true"`
}

var config Config

func init() {
	err := envconfig.Process("app", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	createFirestoreClient()
}
