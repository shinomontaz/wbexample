package common

type Message struct {
	T int `json: "type"`
}

type MessageAddPoint struct {
	Lat  float64 `json: "lat"`
	Long float64 `json: "long"`
}

type MessageGenerateFleet struct {
	//	area float64 `json: "lat"`
}

type MessagePause struct {
	//	area float64 `json: "lat"`
}