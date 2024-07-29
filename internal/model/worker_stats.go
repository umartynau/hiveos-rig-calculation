package model

type WorkerStatsResponse struct {
	Data []WorkerStatsData `json:"data"`
}

type WorkerStatsData struct {
	Time      int64      `json:"time"`
	Units     int        `json:"units"`
	Temp      []int      `json:"temp"`
	Fan       []int      `json:"fan"`
	Power     int        `json:"power"`
	Hashrates []Hashrate `json:"hashrates"`
	Fanflap   struct {
		Fan []int `json:"fan"`
	} `json:"fanflap"`
	Powermeter struct {
		Power       [][]float64 `json:"power"`
		PowerTotal  []float64   `json:"power_total"`
		EnergyTotal []float64   `json:"energy_total"`
	} `json:"powermeter"`
}

type Hashrate struct {
	Algo   string    `json:"algo"`
	Values []float64 `json:"values"`
}
