CREATE TABLE `demo`
(
    `id`           varchar(20) NOT NULL,
    `name`         varchar(255)         DEFAULT NULL,
    `number`       int(11) DEFAULT '0',
    `created_time` datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time` datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;