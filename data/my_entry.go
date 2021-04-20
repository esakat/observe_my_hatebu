package data

type MyEntry struct {
	EntryID         string `json:"eid"`
	URL             string `json:"url"`
	ThreadTimestamp string `json:"thread_timestamp"`
	UpdateTimestamp string `json:"update_timestamp"`
}