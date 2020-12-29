package storage

import "time"

type JsonReport struct {
	Data struct {
		DeletedFiles         int         `json:"DeletedFiles"`
		DeletedFolders       int         `json:"DeletedFolders"`
		ModifiedFiles        int         `json:"ModifiedFiles"`
		ExaminedFiles        int         `json:"ExaminedFiles"`
		OpenedFiles          int         `json:"OpenedFiles"`
		AddedFiles           int         `json:"AddedFiles"`
		SizeOfModifiedFiles  int         `json:"SizeOfModifiedFiles"`
		SizeOfAddedFiles     int         `json:"SizeOfAddedFiles"`
		SizeOfExaminedFiles  int         `json:"SizeOfExaminedFiles"`
		SizeOfOpenedFiles    int         `json:"SizeOfOpenedFiles"`
		NotProcessedFiles    int         `json:"NotProcessedFiles"`
		AddedFolders         int         `json:"AddedFolders"`
		TooLargeFiles        int         `json:"TooLargeFiles"`
		FilesWithError       int         `json:"FilesWithError"`
		ModifiedFolders      int         `json:"ModifiedFolders"`
		ModifiedSymlinks     int         `json:"ModifiedSymlinks"`
		AddedSymlinks        int         `json:"AddedSymlinks"`
		DeletedSymlinks      int         `json:"DeletedSymlinks"`
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
