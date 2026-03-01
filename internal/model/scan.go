package model

import "time"

type ScanStatus string

const (
	ScanStatusPending  ScanStatus = "pending"
	ScanStatusRunning  ScanStatus = "running"
	ScanStatusDone     ScanStatus = "done"
	ScanStatusFailed   ScanStatus = "failed"
)

type Scan struct {
	ID        string     `json:"id"`
	URL       string     `json:"url"`
	Status    ScanStatus `json:"status"`
	Results   any        `json:"results,omitempty"`
	Error     string     `json:"error,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
