package storage

import "time"

type JsonReport struct {
	Data struct {
		DeletedFiles         uint32      `json:"DeletedFiles"`
		DeletedFolders       uint32      `json:"DeletedFolders"`
		ModifiedFiles        uint32      `json:"ModifiedFiles"`
		ExaminedFiles        uint32      `json:"ExaminedFiles"`
		OpenedFiles          uint32      `json:"OpenedFiles"`
		AddedFiles           uint32      `json:"AddedFiles"`
		SizeOfModifiedFiles  uint32      `json:"SizeOfModifiedFiles"`
		SizeOfAddedFiles     uint32      `json:"SizeOfAddedFiles"`
		SizeOfExaminedFiles  uint32      `json:"SizeOfExaminedFiles"`
		SizeOfOpenedFiles    uint32      `json:"SizeOfOpenedFiles"`
		NotProcessedFiles    uint32      `json:"NotProcessedFiles"`
		AddedFolders         uint32      `json:"AddedFolders"`
		TooLargeFiles        uint32      `json:"TooLargeFiles"`
		FilesWithError       uint32      `json:"FilesWithError"`
		ModifiedFolders      uint32      `json:"ModifiedFolders"`
		ModifiedSymlinks     uint32      `json:"ModifiedSymlinks"`
		AddedSymlinks        uint32      `json:"AddedSymlinks"`
		DeletedSymlinks      uint32      `json:"DeletedSymlinks"`
		PartialBackup        bool        `json:"PartialBackup"`
		Dryrun               bool        `json:"Dryrun"`
		MainOperation        string      `json:"MainOperation"`
		CompactResults       interface{} `json:"CompactResults"`
		VacuumResults        interface{} `json:"VacuumResults"`
		DeleteResults        interface{} `json:"DeleteResults"`
		RepairResults        interface{} `json:"RepairResults"`
		TestResults          TestResults
		ParsedResult         string    `json:"ParsedResult"`
		Version              string    `json:"Version"`
		EndTime              time.Time `json:"EndTime"`
		BeginTime            time.Time `json:"BeginTime"`
		Duration             string    `json:"Duration"`
		MessagesActualLength int       `json:"MessagesActualLength"`
		WarningsActualLength int       `json:"WarningsActualLength"`
		ErrorsActualLength   int       `json:"ErrorsActualLength"`
		BackendStatistics    BackendStatistics
	} `json:"Data"`
	Extra    Extra
	LogLines []interface{} `json:"LogLines"`
}
