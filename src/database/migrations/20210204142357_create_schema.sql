-- +goose Up
-- +goose StatementBegin
create  schema accounts default  character set utf8mb4
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop schema accounts
-- +goose StatementEnd
