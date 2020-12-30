// Package duplicati provides processing of Duplicati send-http json data to the reports table in the MySQL databas.
//TODO: Added key based ACL
package duplicati

import (
	"./storage"
	"database/sql"
	"fmt"
)

var (
	// table is the table name.
	table = "duplicati_status_reports"
)

// Connection is an interface for making queries.
type Connection interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

// Create adds an item.
func Create(db Connection, report storage.DuplicatiStatusReports) (sql.Result, error) {
	sql := fmt.Sprintf(`
		INSERT INTO %v
		(
			server,
			operation_name,
			backup_name,
			deleted_files,
			deleted_folders,
			modified_files,
			examined_files,
			opened_files,
			added_files,
			size_of_modified_files,
			size_of_added_files,
			sizes_of_examined_files,
			size_of_opened_files,
			not_processed_files,
			added_folders,
			too_large_files,
			files_with_error,
			modified_folders,
			modified_symlinks,
			added_symlinks,
			deleted_symlinks,
			partial_backup,
			dryrun,
			main_operation,
			parsed_result,
			version,
			end_time,
			begin_time,
			duration)
		VALUES
        ('%s','%s','%s',%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%t,%t,'%s','%s','%s','%s','%s','%s')
		`, table,
		report.Server,
		report.OperationName,
		report.BackupName,
		report.DeletedFiles,
		report.DeletedFolders,
		report.ModifiedFiles,
		report.ExaminedFiles,
		report.OpenedFiles,
		report.AddedFiles,
		report.SizeOfExaminedFiles,
		report.SizeOfAddedFiles,
		report.SizeOfExaminedFiles,
		report.SizeOfOpenedFiles,
		report.NotProcessedFiles,
		report.AddedFolders,
		report.TooLargeFiles,
		report.FilesWithError,
		report.ModifiedFolders,
		report.ModifiedSymlinks,
		report.AddedSymlinks,
		report.DeletedSymlinks,
		report.PartialBackup,
		report.Dryrun,
		report.MainOperation,
		report.ParsedResult,
		report.Version,
		report.EndTime,
		report.BeginTime,
		report.Duration)

	//log.Println(sql)

	result, err := db.Exec(sql)
	return result, err
}

func CreateSummary(server string, report storage.JsonReport) (summaryReport storage.DuplicatiStatusReports) {
	//summary.ID = ''
	summaryReport.Server = server
	summaryReport.OperationName = report.Extra.OperationName
	summaryReport.BackupName = report.Extra.BackupName
	summaryReport.DeletedFiles = report.Data.DeletedFiles
	summaryReport.DeletedFolders = report.Data.DeletedFolders
	summaryReport.ModifiedFiles = report.Data.ModifiedFiles
	summaryReport.ExaminedFiles = report.Data.ExaminedFiles
	summaryReport.AddedFiles = report.Data.AddedFiles
	summaryReport.SizeOfModifiedFiles = report.Data.SizeOfModifiedFiles
	summaryReport.SizeOfAddedFiles = report.Data.SizeOfAddedFiles
	summaryReport.SizeOfExaminedFiles = report.Data.SizeOfExaminedFiles
	summaryReport.SizeOfOpenedFiles = report.Data.SizeOfOpenedFiles
	summaryReport.NotProcessedFiles = report.Data.NotProcessedFiles
	summaryReport.AddedFolders = report.Data.AddedFolders
	summaryReport.TooLargeFiles = report.Data.TooLargeFiles
	summaryReport.FilesWithError = report.Data.FilesWithError
	summaryReport.ModifiedFolders = report.Data.ModifiedFolders
	summaryReport.ModifiedSymlinks = report.Data.ModifiedSymlinks
	summaryReport.AddedSymlinks = report.Data.AddedSymlinks
	summaryReport.DeletedSymlinks = report.Data.DeletedSymlinks
	summaryReport.PartialBackup = report.Data.PartialBackup
	summaryReport.Dryrun = report.Data.Dryrun
	summaryReport.MainOperation = report.Data.MainOperation
	summaryReport.ParsedResult = report.Data.ParsedResult
	summaryReport.Version = report.Data.Version
	summaryReport.EndTime = report.Data.EndTime
	summaryReport.BeginTime = report.Data.BeginTime
	summaryReport.Duration = report.Data.Duration

	return summaryReport
}
