CREATE TABLE `comments` (
`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
`page_id` int(11) NOT NULL,
`comment_guid` varchar(256) DEFAULT NULL,
`comment_name` varchar(64) DEFAULT NULL,
`comment_email` varchar(128) DEFAULT NULL,
`comment_text` mediumtext,
`comment_date` timestamp NULL DEFAULT NULL,
PRIMARY KEY (`id`),
KEY `page_id` (`page_id`)
) ENGINE=InnoDB;