// Package duplicati provides processing of Duplicati send-http json data to the reports table in the MySQL databas.
//TODO: Added key based ACL
package duplicati

import (
	"database/sql"
	"time"
)

var (
	// table is the table name.
	table = "reports"
)

type JsonReportOld struct {
	Data struct {
		DeletedFiles        int         `json:"DeletedFiles"`
		DeletedFolders      int         `json:"DeletedFolders"`
		ModifiedFiles       int         `json:"ModifiedFiles"`
		ExaminedFiles       int         `json:"ExaminedFiles"`
		OpenedFiles         int         `json:"OpenedFiles"`
		AddedFiles          int         `json:"AddedFiles"`
		SizeOfModifiedFiles int         `json:"SizeOfModifiedFiles"`
		SizeOfAddedFiles    int         `json:"SizeOfAddedFiles"`
		SizeOfExaminedFiles int         `json:"SizeOfExaminedFiles"`
		SizeOfOpenedFiles   int         `json:"SizeOfOpenedFiles"`
		NotProcessedFiles   int         `json:"NotProcessedFiles"`
		AddedFolders        int         `json:"AddedFolders"`
		TooLargeFiles       int         `json:"TooLargeFiles"`
		FilesWithError      int         `json:"FilesWithError"`
		ModifiedFolders     int         `json:"ModifiedFolders"`
		ModifiedSymlinks    int         `json:"ModifiedSymlinks"`
		AddedSymlinks       int         `json:"AddedSymlinks"`
		DeletedSymlinks     int         `json:"DeletedSymlinks"`
		PartialBackup       bool        `json:"PartialBackup"`
		Dryrun              bool        `json:"Dryrun"`
		MainOperation       string      `json:"MainOperation"`
		CompactResults      interface{} `json:"CompactResults"`
		VacuumResults       interface{} `json:"VacuumResults"`
		DeleteResults       interface{} `json:"DeleteResults"`
		RepairResults       interface{} `json:"RepairResults"`
		TestResults         struct {
			MainOperation             string `json:"MainOperation"`
			VerificationsActualLength int    `json:"VerificationsActualLength"`
			Verifications             []struct {
				Key   string        `json:"Key"`
				Value []interface{} `json:"Value"`
			} `json:"Verifications"`
			ParsedResult         string    `json:"ParsedResult"`
			Version              string    `json:"Version"`
			EndTime              time.Time `json:"EndTime"`
			BeginTime            time.Time `json:"BeginTime"`
			Duration             string    `json:"Duration"`
			MessagesActualLength int       `json:"MessagesActualLength"`
			WarningsActualLength int       `json:"WarningsActualLength"`
			ErrorsActualLength   int       `json:"ErrorsActualLength"`
			BackendStatistics    struct {
				RemoteCalls          int       `json:"RemoteCalls"`
				BytesUploaded        int       `json:"BytesUploaded"`
				BytesDownloaded      int       `json:"BytesDownloaded"`
				FilesUploaded        int       `json:"FilesUploaded"`
				FilesDownloaded      int       `json:"FilesDownloaded"`
				FilesDeleted         int       `json:"FilesDeleted"`
				FoldersCreated       int       `json:"FoldersCreated"`
				RetryAttempts        int       `json:"RetryAttempts"`
				UnknownFileSize      int       `json:"UnknownFileSize"`
				UnknownFileCount     int       `json:"UnknownFileCount"`
				KnownFileCount       int       `json:"KnownFileCount"`
				KnownFileSize        int       `json:"KnownFileSize"`
				LastBackupDate       string    `json:"LastBackupDate"`
				BackupListCount      int       `json:"BackupListCount"`
				TotalQuotaSpace      int64     `json:"TotalQuotaSpace"`
				FreeQuotaSpace       int64     `json:"FreeQuotaSpace"`
				AssignedQuotaSpace   int       `json:"AssignedQuotaSpace"`
				ReportedQuotaError   bool      `json:"ReportedQuotaError"`
				ReportedQuotaWarning bool      `json:"ReportedQuotaWarning"`
				MainOperation        string    `json:"MainOperation"`
				ParsedResult         string    `json:"ParsedResult"`
				Version              string    `json:"Version"`
				EndTime              string    `json:"EndTime"`
				BeginTime            time.Time `json:"BeginTime"`
				Duration             string    `json:"Duration"`
				MessagesActualLength int       `json:"MessagesActualLength"`
				WarningsActualLength int       `json:"WarningsActualLength"`
				ErrorsActualLength   int       `json:"ErrorsActualLength"`
			} `json:"BackendStatistics"`
		} `json:"TestResults"`
		ParsedResult         string    `json:"ParsedResult"`
		Version              string    `json:"Version"`
		EndTime              time.Time `json:"EndTime"`
		BeginTime            time.Time `json:"BeginTime"`
		Duration             string    `json:"Duration"`
		MessagesActualLength int       `json:"MessagesActualLength"`
		WarningsActualLength int       `json:"WarningsActualLength"`
		ErrorsActualLength   int       `json:"ErrorsActualLength"`
		BackendStatistics    struct {
			RemoteCalls          int       `json:"RemoteCalls"`
			BytesUploaded        int       `json:"BytesUploaded"`
			BytesDownloaded      int       `json:"BytesDownloaded"`
			FilesUploaded        int       `json:"FilesUploaded"`
			FilesDownloaded      int       `json:"FilesDownloaded"`
			FilesDeleted         int       `json:"FilesDeleted"`
			FoldersCreated       int       `json:"FoldersCreated"`
			RetryAttempts        int       `json:"RetryAttempts"`
			UnknownFileSize      int       `json:"UnknownFileSize"`
			UnknownFileCount     int       `json:"UnknownFileCount"`
			KnownFileCount       int       `json:"KnownFileCount"`
			KnownFileSize        int       `json:"KnownFileSize"`
			LastBackupDate       string    `json:"LastBackupDate"`
			BackupListCount      int       `json:"BackupListCount"`
			TotalQuotaSpace      int64     `json:"TotalQuotaSpace"`
			FreeQuotaSpace       int64     `json:"FreeQuotaSpace"`
			AssignedQuotaSpace   int       `json:"AssignedQuotaSpace"`
			ReportedQuotaError   bool      `json:"ReportedQuotaError"`
			ReportedQuotaWarning bool      `json:"ReportedQuotaWarning"`
			MainOperation        string    `json:"MainOperation"`
			ParsedResult         string    `json:"ParsedResult"`
			Version              string    `json:"Version"`
			EndTime              string    `json:"EndTime"`
			BeginTime            time.Time `json:"BeginTime"`
			Duration             string    `json:"Duration"`
			MessagesActualLength int       `json:"MessagesActualLength"`
			WarningsActualLength int       `json:"WarningsActualLength"`
			ErrorsActualLength   int       `json:"ErrorsActualLength"`
		} `json:"BackendStatistics"`
	} `json:"Data"`
	Extra struct {
		OperationName string `json:"OperationName"`
		BackupName    string `json:"backup-name"`
	} `json:"Extra"`
	LogLines []interface{} `json:"LogLines"`
}

// Connection is an interface for making queries.
type Connection interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}
