package api

type CEXResponse struct {
	Timestamp   string  `json:"timestamp"`
	Low         string  `json:"low"`
	High        string  `json:"high"`
	Last        string  `json:"last"`
	Volume      string  `json:"volume"`
	PriceChange string  `json:"priceChange"`
	Bid         float64 `json:"bid"`
}
