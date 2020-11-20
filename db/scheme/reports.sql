CREATE TABLE `reports` (
  `id` int(10) NOT NULL,
  `Server` varchar(255) NOT NULL,
  `Data_DeletedFiles` int(11) DEFAULT NULL,
  `Data_DeletedFolders` int(11) DEFAULT NULL,
  `Data_ModifiedFiles` int(11) DEFAULT NULL,
  `Data_ExaminedFiles` int(11) DEFAULT NULL,
  `Data_OpenedFiles` int(11) DEFAULT NULL,
  `Data_AddedFiles` int(11) DEFAULT NULL,
  `Data_SizeOfModifiedFiles` int(11) DEFAULT NULL,
  `Data_SizeOfAddedFiles` int(11) DEFAULT NULL,
  `Data_SizeOfExaminedFiles` int(11) DEFAULT NULL,
  `Data_SizeOfOpenedFiles` int(11) DEFAULT NULL,
  `Data_NotProcessedFiles` int(11) DEFAULT NULL,
  `Data_AddedFolders` int(11) DEFAULT NULL,
  `Data_TooLargeFiles` int(11) DEFAULT NULL,
  `Data_FilesWithError` int(11) DEFAULT NULL,
  `Data_ModifiedFolders` int(11) DEFAULT NULL,
  `Data_ModifiedSymlinks` int(11) DEFAULT NULL,
  `Data_AddedSymlinks` int(11) DEFAULT NULL,
  `Data_DeletedSymlinks` int(11) DEFAULT NULL,
  `Data_PartialBackup` varchar(5) CHARACTER SET utf8 DEFAULT NULL,
  `Data_Dryrun` varchar(5) CHARACTER SET utf8 DEFAULT NULL,
  `Data_MainOperation` varchar(6) CHARACTER SET utf8 DEFAULT NULL,
  `Data_CompactResults` int(11) DEFAULT NULL,
  `Data_VacuumResults` int(11) DEFAULT NULL,
  `Data_DeleteResults` int(11) DEFAULT NULL,
  `Data_RepairResults` int(11) DEFAULT NULL,
  `Data_TestResults_MainOperation` varchar(4) CHARACTER SET utf8 DEFAULT NULL,
  `Data_TestResults_VerificationsActualLength` int(11) DEFAULT NULL,
  `Data_TestResults_Verifications_Key` varchar(58) CHARACTER SET utf8 DEFAULT NULL,
  `Data_TestResults_ParsedResult` varchar(7) CHARACTER SET utf8 DEFAULT NULL,
  `Data_TestResults_Version` varchar(33) CHARACTER SET utf8 DEFAULT NULL,
  `Data_TestResults_EndTime` varchar(28) CHARACTER SET utf8 DEFAULT NULL,
  `Data_TestResults_BeginTime` varchar(28) CHARACTER SET utf8 DEFAULT NULL,
  `Data_TestResults_Duration` varchar(16) CHARACTER SET utf8 DEFAULT NULL,
  `Data_TestResults_MessagesActualLength` int(11) DEFAULT NULL,
  `Data_TestResults_WarningsActualLength` int(11) DEFAULT NULL,
  `Data_TestResults_ErrorsActualLength` int(11) DEFAULT NULL,
  `Data_TestResults_BackendStatistics_RemoteCalls` int(11) DEFAULT NULL,
  `Data_TestResults_BackendStatistics_BytesUploaded` int(11) DEFAULT NULL,
  `Data_TestResults_BackendStatistics_BytesDownloaded` int(11) DEFAULT NULL,
  `Data_TestResults_BackendStatistics_FilesUploaded` int(11) DEFAULT NULL,
  `Data_TestResults_BackendStatistics_FilesDownloaded` int(11) DEFAULT NULL,
  `Data_TestResults_BackendStatistics_FilesDeleted` int(11) DEFAULT NULL,
  `Data_TestResults_BackendStatistics_FoldersCreated` int(11) DEFAULT NULL,
  `Data_TestResults_BackendStatistics_RetryAttempts` int(11) DEFAULT NULL,
  `Data_TestResults_BackendStatistics_UnknownFileSize` int(11) DEFAULT NULL,
  `Data_TestResults_BackendStatistics_UnknownFileCount` int(11) DEFAULT NULL,
  `Data_TestResults_BackendStatistics_KnownFileCount` int(11) DEFAULT NULL,
  `Data_TestResults_BackendStatistics_KnownFileSize` int(11) DEFAULT NULL,
  `Data_TestResults_BackendStatistics_LastBackupDate` varchar(25) CHARACTER SET utf8 DEFAULT NULL,
  `Data_TestResults_BackendStatistics_BackupListCount` int(11) DEFAULT NULL,
  `Data_TestResults_BackendStatistics_TotalQuotaSpace` bigint(20) DEFAULT NULL,
  `Data_TestResults_BackendStatistics_FreeQuotaSpace` bigint(20) DEFAULT NULL,
  `Data_TestResults_BackendStatistics_AssignedQuotaSpace` int(11) DEFAULT NULL,
  `Data_TestResults_BackendStatistics_ReportedQuotaError` varchar(5) CHARACTER SET utf8 DEFAULT NULL,
  `Data_TestResults_BackendStatistics_ReportedQuotaWarning` varchar(5) CHARACTER SET utf8 DEFAULT NULL,
  `Data_TestResults_BackendStatistics_MainOperation` varchar(6) CHARACTER SET utf8 DEFAULT NULL,
  `Data_TestResults_BackendStatistics_ParsedResult` varchar(7) CHARACTER SET utf8 DEFAULT NULL,
  `Data_TestResults_BackendStatistics_Version` varchar(33) CHARACTER SET utf8 DEFAULT NULL,
  `Data_TestResults_BackendStatistics_EndTime` varchar(19) CHARACTER SET utf8 DEFAULT NULL,
  `Data_TestResults_BackendStatistics_BeginTime` varchar(28) CHARACTER SET utf8 DEFAULT NULL,
  `Data_TestResults_BackendStatistics_Duration` varchar(8) CHARACTER SET utf8 DEFAULT NULL,
  `Data_TestResults_BackendStatistics_MessagesActualLength` int(11) DEFAULT NULL,
  `Data_TestResults_BackendStatistics_WarningsActualLength` int(11) DEFAULT NULL,
  `Data_TestResults_BackendStatistics_ErrorsActualLength` int(11) DEFAULT NULL,
  `Data_ParsedResult` varchar(7) CHARACTER SET utf8 DEFAULT NULL,
  `Data_Version` varchar(33) CHARACTER SET utf8 DEFAULT NULL,
  `Data_EndTime` varchar(28) CHARACTER SET utf8 DEFAULT NULL,
  `Data_BeginTime` varchar(28) CHARACTER SET utf8 DEFAULT NULL,
  `Data_Duration` varchar(16) CHARACTER SET utf8 DEFAULT NULL,
  `Data_MessagesActualLength` int(11) DEFAULT NULL,
  `Data_WarningsActualLength` int(11) DEFAULT NULL,
  `Data_ErrorsActualLength` int(11) DEFAULT NULL,
  `Data_BackendStatistics_RemoteCalls` int(11) DEFAULT NULL,
  `Data_BackendStatistics_BytesUploaded` int(11) DEFAULT NULL,
  `Data_BackendStatistics_BytesDownloaded` int(11) DEFAULT NULL,
  `Data_BackendStatistics_FilesUploaded` int(11) DEFAULT NULL,
  `Data_BackendStatistics_FilesDownloaded` int(11) DEFAULT NULL,
  `Data_BackendStatistics_FilesDeleted` int(11) DEFAULT NULL,
  `Data_BackendStatistics_FoldersCreated` int(11) DEFAULT NULL,
  `Data_BackendStatistics_RetryAttempts` int(11) DEFAULT NULL,
  `Data_BackendStatistics_UnknownFileSize` int(11) DEFAULT NULL,
  `Data_BackendStatistics_UnknownFileCount` int(11) DEFAULT NULL,
  `Data_BackendStatistics_KnownFileCount` int(11) DEFAULT NULL,
  `Data_BackendStatistics_KnownFileSize` int(11) DEFAULT NULL,
  `Data_BackendStatistics_LastBackupDate` varchar(25) CHARACTER SET utf8 DEFAULT NULL,
  `Data_BackendStatistics_BackupListCount` int(11) DEFAULT NULL,
  `Data_BackendStatistics_TotalQuotaSpace` bigint(20) DEFAULT NULL,
  `Data_BackendStatistics_FreeQuotaSpace` bigint(20) DEFAULT NULL,
  `Data_BackendStatistics_AssignedQuotaSpace` int(11) DEFAULT NULL,
  `Data_BackendStatistics_ReportedQuotaError` varchar(5) CHARACTER SET utf8 DEFAULT NULL,
  `Data_BackendStatistics_ReportedQuotaWarning` varchar(5) CHARACTER SET utf8 DEFAULT NULL,
  `Data_BackendStatistics_MainOperation` varchar(6) CHARACTER SET utf8 DEFAULT NULL,
  `Data_BackendStatistics_ParsedResult` varchar(7) CHARACTER SET utf8 DEFAULT NULL,
  `Data_BackendStatistics_Version` varchar(33) CHARACTER SET utf8 DEFAULT NULL,
  `Data_BackendStatistics_EndTime` varchar(19) CHARACTER SET utf8 DEFAULT NULL,
  `Data_BackendStatistics_BeginTime` varchar(28) CHARACTER SET utf8 DEFAULT NULL,
  `Data_BackendStatistics_Duration` varchar(8) CHARACTER SET utf8 DEFAULT NULL,
  `Data_BackendStatistics_MessagesActualLength` int(11) DEFAULT NULL,
  `Data_BackendStatistics_WarningsActualLength` int(11) DEFAULT NULL,
  `Data_BackendStatistics_ErrorsActualLength` int(11) DEFAULT NULL,
  `Extra_OperationName` varchar(6) CHARACTER SET utf8 DEFAULT NULL,
  `Extra_backup_name` varchar(19) CHARACTER SET utf8 DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `report`
--

INSERT INTO `reports` (`id`, `Server`, `Data_DeletedFiles`, `Data_DeletedFolders`, `Data_ModifiedFiles`, `Data_ExaminedFiles`, `Data_OpenedFiles`, `Data_AddedFiles`, `Data_SizeOfModifiedFiles`, `Data_SizeOfAddedFiles`, `Data_SizeOfExaminedFiles`, `Data_SizeOfOpenedFiles`, `Data_NotProcessedFiles`, `Data_AddedFolders`, `Data_TooLargeFiles`, `Data_FilesWithError`, `Data_ModifiedFolders`, `Data_ModifiedSymlinks`, `Data_AddedSymlinks`, `Data_DeletedSymlinks`, `Data_PartialBackup`, `Data_Dryrun`, `Data_MainOperation`, `Data_CompactResults`, `Data_VacuumResults`, `Data_DeleteResults`, `Data_RepairResults`, `Data_TestResults_MainOperation`, `Data_TestResults_VerificationsActualLength`, `Data_TestResults_Verifications_Key`, `Data_TestResults_ParsedResult`, `Data_TestResults_Version`, `Data_TestResults_EndTime`, `Data_TestResults_BeginTime`, `Data_TestResults_Duration`, `Data_TestResults_MessagesActualLength`, `Data_TestResults_WarningsActualLength`, `Data_TestResults_ErrorsActualLength`, `Data_TestResults_BackendStatistics_RemoteCalls`, `Data_TestResults_BackendStatistics_BytesUploaded`, `Data_TestResults_BackendStatistics_BytesDownloaded`, `Data_TestResults_BackendStatistics_FilesUploaded`, `Data_TestResults_BackendStatistics_FilesDownloaded`, `Data_TestResults_BackendStatistics_FilesDeleted`, `Data_TestResults_BackendStatistics_FoldersCreated`, `Data_TestResults_BackendStatistics_RetryAttempts`, `Data_TestResults_BackendStatistics_UnknownFileSize`, `Data_TestResults_BackendStatistics_UnknownFileCount`, `Data_TestResults_BackendStatistics_KnownFileCount`, `Data_TestResults_BackendStatistics_KnownFileSize`, `Data_TestResults_BackendStatistics_LastBackupDate`, `Data_TestResults_BackendStatistics_BackupListCount`, `Data_TestResults_BackendStatistics_TotalQuotaSpace`, `Data_TestResults_BackendStatistics_FreeQuotaSpace`, `Data_TestResults_BackendStatistics_AssignedQuotaSpace`, `Data_TestResults_BackendStatistics_ReportedQuotaError`, `Data_TestResults_BackendStatistics_ReportedQuotaWarning`, `Data_TestResults_BackendStatistics_MainOperation`, `Data_TestResults_BackendStatistics_ParsedResult`, `Data_TestResults_BackendStatistics_Version`, `Data_TestResults_BackendStatistics_EndTime`, `Data_TestResults_BackendStatistics_BeginTime`, `Data_TestResults_BackendStatistics_Duration`, `Data_TestResults_BackendStatistics_MessagesActualLength`, `Data_TestResults_BackendStatistics_WarningsActualLength`, `Data_TestResults_BackendStatistics_ErrorsActualLength`, `Data_ParsedResult`, `Data_Version`, `Data_EndTime`, `Data_BeginTime`, `Data_Duration`, `Data_MessagesActualLength`, `Data_WarningsActualLength`, `Data_ErrorsActualLength`, `Data_BackendStatistics_RemoteCalls`, `Data_BackendStatistics_BytesUploaded`, `Data_BackendStatistics_BytesDownloaded`, `Data_BackendStatistics_FilesUploaded`, `Data_BackendStatistics_FilesDownloaded`, `Data_BackendStatistics_FilesDeleted`, `Data_BackendStatistics_FoldersCreated`, `Data_BackendStatistics_RetryAttempts`, `Data_BackendStatistics_UnknownFileSize`, `Data_BackendStatistics_UnknownFileCount`, `Data_BackendStatistics_KnownFileCount`, `Data_BackendStatistics_KnownFileSize`, `Data_BackendStatistics_LastBackupDate`, `Data_BackendStatistics_BackupListCount`, `Data_BackendStatistics_TotalQuotaSpace`, `Data_BackendStatistics_FreeQuotaSpace`, `Data_BackendStatistics_AssignedQuotaSpace`, `Data_BackendStatistics_ReportedQuotaError`, `Data_BackendStatistics_ReportedQuotaWarning`, `Data_BackendStatistics_MainOperation`, `Data_BackendStatistics_ParsedResult`, `Data_BackendStatistics_Version`, `Data_BackendStatistics_EndTime`, `Data_BackendStatistics_BeginTime`, `Data_BackendStatistics_Duration`, `Data_BackendStatistics_MessagesActualLength`, `Data_BackendStatistics_WarningsActualLength`, `Data_BackendStatistics_ErrorsActualLength`, `Extra_OperationName`, `Extra_backup_name`) VALUES
(1, '', 0, 0, 0, 20, 0, 0, 0, 0, 13289479, 0, 0, 0, 0, 0, 0, 0, 0, 0, 'False', 'False', 'Backup', NULL, NULL, NULL, NULL, 'Test', NULL, 'duplicati-b11f4b10e854b44ae93e2783ccce405cb.dblock.zip.aes', 'Success', '2.0.5.1 (2.0.5.1_beta_2020-01-18)', '2020-11-17T20:29:41.2835716Z', '2020-11-17T20:29:40.8147927Z', '00:00:00.4687789', 0, 0, 0, 5, 0, 8246807, 0, 3, 0, 0, 0, 0, 1, 3, 8246807, '2020-11-16T19:39:15-05:00', 1, 3000581881856, 671590731776, -1, 'False', 'False', 'Backup', 'Success', '2.0.5.1 (2.0.5.1_beta_2020-01-18)', '0001-01-01T00:00:00', '2020-11-17T20:29:40.4241435Z', '00:00:00', 0, 0, 0, 'Success', '2.0.5.1 (2.0.5.1_beta_2020-01-18)', '2020-11-17T20:29:41.7836077Z', '2020-11-17T20:29:40.4241435Z', '00:00:01.3594642', 11, 0, 0, 5, 0, 8246807, 0, 3, 0, 0, 0, 0, 1, 3, 8246807, '2020-11-16T19:39:15-05:00', 1, 3000581881856, 671590731776, -1, 'False', 'False', 'Backup', 'Success', '2.0.5.1 (2.0.5.1_beta_2020-01-18)', '0001-01-01T00:00:00', '2020-11-17T20:29:40.4241435Z', '00:00:00', 0, 0, 0, 'Backup', 'Testing Send Status');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `report`
--
ALTER TABLE `reports`
  ADD PRIMARY KEY (`id`),
  ADD KEY `Server` (`Server`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `report`
--
ALTER TABLE `reports`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
COMMIT;
