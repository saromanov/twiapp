package structs

// Message provides basic component for communication between Websockets
// and RabbitMQ
type Message struct {
	Event string
	Data  string
}

// Tweet provides main item for each tweet
type Tweet struct {
	Author string `json:"author"`
	Text string   `json:"text"`
}