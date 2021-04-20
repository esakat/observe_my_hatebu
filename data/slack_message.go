package data

type SlackMessage struct {
	ThreadTimestamp string `json:"thread_timestamp"`
	Text            string `json:"text"`
	UserName        string `json:"username"`
	IconURL         string `json:"icon_url"`
	ChannelID       string `json:"channel_id"`
}
