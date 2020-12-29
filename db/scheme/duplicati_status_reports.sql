--
-- Table structure for table `duplicati_status_reports`
--

CREATE TABLE `duplicati_status_reports` (
  `id` int(10) NOT NULL,
  `server` varchar(255) NOT NULL,
  `deleted_files` int(10) NOT NULL,
  `deleted_folders` int(10) NOT NULL,
  `modified_files` int(10) NOT NULL,
  `examined_files` int(10) NOT NULL,
  `opened_files` int(10) NOT NULL,
  `added_files` int(10) NOT NULL,
  `size_of_modified_files` int(10) NOT NULL,
  `size_of_added_files` int(10) NOT NULL,
  `sizes_of_examined_files` int(10) NOT NULL,
  `size_of_opened_files` int(10) NOT NULL,
  `not_processed_files` int(10) NOT NULL,
  `added_folders` int(10) NOT NULL,
  `too_large_files` int(10) NOT NULL,
  `files_with_error` int(10) NOT NULL,
  `modified_folders` int(10) NOT NULL,
  `modified_symlinks` int(10) NOT NULL,
  `added_symlinks` int(10) NOT NULL,
  `deleted_symlinks` int(10) NOT NULL,
  `partial_backup` tinyint(1) NOT NULL,
  `dryrun` tinyint(1) NOT NULL,
  `main_operation` varchar(255) NOT NULL,
  `parsed_result` varchar(255) NOT NULL,
  `version` varchar(255) NOT NULL,
  `end_time` datetime NOT NULL,
  `begin_time` datetime NOT NULL,
  `duration` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `duplicati_status_reports`
--
ALTER TABLE `duplicati_status_reports`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `duplicati_status_reports`
--
ALTER TABLE `duplicati_status_reports`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT;
COMMIT;