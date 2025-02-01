-- +goose Up
-- +goose StatementBegin
create table urls
(
    short_url varchar(10) primary key,
    long_url  text        not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table urls
-- +goose StatementEnd
