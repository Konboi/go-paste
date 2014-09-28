DROP TABLE `np`;
CREATE TABLE IF NOT EXISTS `np` (
 `code` varchar(255) NOT NULL UNIQUE,
 `title` varchar(255),
 `body` text,
 `created_at` datetime default NOW()
) DEFAULT CHARSET=utf8;
