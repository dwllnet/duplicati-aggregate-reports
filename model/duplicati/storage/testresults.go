package storage

import "time"

type TestResults struct {
	MainOperation             string `json:"MainOperation"`
	VerificationsActualLength int    `json:"VerificationsActualLength"`
	Verifications             []Verifications
	ParsedResult              string    `json:"ParsedResult"`
	Version                   string    `json:"Version"`
	EndTime                   time.Time `json:"EndTime"`
	BeginTime                 time.Time `json:"BeginTime"`
	Duration                  string    `json:"Duration"`
	MessagesActualLength      int       `json:"MessagesActualLength"`
	WarningsActualLength      int       `json:"WarningsActualLength"`
	ErrorsActualLength        int       `json:"ErrorsActualLength"`
	BackendStatistics         BackendStatistics
}
