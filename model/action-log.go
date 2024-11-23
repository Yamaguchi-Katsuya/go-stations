package model

import "time"

type ActionLog struct {
	Timestamp time.Time `json:"timestamp"`
	Latency   int64     `json:"latency"`
	Path      string    `json:"path"`
	OS        string    `json:"os"`
}

type ActionLogResponse struct{}
