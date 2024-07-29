package model

type Report struct {
	FarmName          string `json:"farm_name"`
	FarmId            int    `json:"farm_id"`
	WorkerName        string `json:"worker_name"`
	WorkerId          int    `json:"worker_id"`
	WorkerDescription string `json:"worker_description"`
	UptimeHours       int    `json:"uptime_hours"`
}
