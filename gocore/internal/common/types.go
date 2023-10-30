package common

type Message struct {
	Type int `json: "type"`
}

type MessageAddPoint struct {
	Lat  float64 `json: "lat"`
	Long float64 `json: "long"`
}

type MessageGenerateFleet struct {
	//	area float64 `json: "lat"`
	Type int  `json:"type"`
	Num  int  `json:"num"`
	Area Area `json:"area"`
}

type MessagePauseFleet struct {
	//	area float64 `json: "lat"`
}
type MessageStartFleet struct {
	//	area float64 `json: "lat"`
}

type Area struct {
	Min []float64 `json:"min"`
	Max []float64 `json:"max"`
}
type Point struct {
	Id   int     `json: id`
	Lat  float64 `json: "lat"`
	Long float64 `json: "long"`
}

type Coord struct {
	Lat  float64 `json: "lat"`
	Long float64 `json: "long"`
}

type Truck struct {
	Id             int `json: id`
	Coord          `json: "coord"`
	Destination    Coord
	HasDestination bool
}

type Position struct {
	TruckId int `json: truck_id`
	Coord   `json: "coord"`
	Prev    Coord `json: "prev"`
}
