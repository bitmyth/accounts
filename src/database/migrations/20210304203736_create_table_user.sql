-- +goose Up
-- +goose StatementBegin
CREATE TABLE `users` (
                       `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                       `name` varchar(100) DEFAULT NULL,
                       `password` varchar(100) DEFAULT NULL,
                       `email` varchar(100),
                       `phone` varchar(191) DEFAULT NULL,
                       `avatar` varchar(100) DEFAULT NULL,
                       `created_at` datetime(3) DEFAULT NULL,
                       `updated_at` datetime(3) DEFAULT NULL,
                       `deleted_at` datetime(3) DEFAULT NULL,
                       PRIMARY KEY (`id`),
                       KEY `idx_users_name` (`name`),
                       KEY `idx_users_phone` (`phone`),
                       KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE uesrs;
-- +goose StatementEnd
