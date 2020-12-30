package storage

import (
	"time"
)

type DuplicatiStatusReports struct {
	ID                  uint32    `db:"id"`
	Server              string    `db:"server"`
	OperationName       string    `db:"operation_name"`
	BackupName          string    `db:"backup_name"`
	DeletedFiles        uint32    `db:"deleted_files"`
	DeletedFolders      uint32    `db:"deleted_folders"`
	ModifiedFiles       uint32    `db:"modified_files"`
	ExaminedFiles       uint32    `db:"examined_files"`
	OpenedFiles         uint32    `db:"opened_files"`
	AddedFiles          uint32    `db:"added_files"`
	SizeOfModifiedFiles uint32    `db:"size_of_modified_files"`
	SizeOfAddedFiles    uint32    `db:"size_of_added_files"`
	SizeOfExaminedFiles uint32    `db:"sizes_of_examined_files"`
	SizeOfOpenedFiles   uint32    `db:"size_of_opened_files"`
	NotProcessedFiles   uint32    `db:"not_processed_files"`
	AddedFolders        uint32    `db:"added_folders"`
	TooLargeFiles       uint32    `db:"too_large_files"`
	FilesWithError      uint32    `db:"files_with_error"`
	ModifiedFolders     uint32    `db:"modified_folders"`
	ModifiedSymlinks    uint32    `db:"modified_symlinks"`
	AddedSymlinks       uint32    `db:"added_symlinks"`
	DeletedSymlinks     uint32    `db:"deleted_symlinks"`
	PartialBackup       bool      `db:"partial_backup"`
	Dryrun              bool      `db:"dryrun"`
	MainOperation       string    `db:"main_operation"`
	ParsedResult        string    `db:"parsed_result"`
	Version             string    `db:"version"`
	EndTime             time.Time `db:"end_time"`
	BeginTime           time.Time `db:"begin_time"`
	Duration            string    `db:"duration"`
}
