-- +goose Up
-- +goose StatementBegin
create table urls
(
    short_url text primary key,
    long_url  text not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table urls;
-- +goose StatementEnd
