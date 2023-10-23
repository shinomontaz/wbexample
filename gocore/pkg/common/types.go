package common

type Message struct {
	t    int    `json: "type"`
	data string `json: "data"`
}
