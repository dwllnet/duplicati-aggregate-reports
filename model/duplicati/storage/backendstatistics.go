package storage

import "time"

type BackendStatistics struct {
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
}
