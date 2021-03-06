package models

// Resource :nodoc:
type Resource struct {
	UUID      string `json:"uuid"`
	Comodity  string `json:"komoditas"`
	Province  string `json:"area_provinsi"`
	City      string `json:"area_kota"`
	Size      string `json:"size"`
	PriceIDR  string `json:"price"`
	PriceUSD  string `json:"price_usd"`
	Date      string `json:"tgl_parsed"`
	Timestamp string `json:"timestamp"`
}
