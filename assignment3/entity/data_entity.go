package entity

type Data struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type DataStatus struct {
	WaterStatus string `json:"waterStatus"`
	WindStatus  string `json:"windStatus"`
}
