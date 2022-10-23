-- +goose Up
create table notes 
(
    id bigserial primary key,
    title text not null,
    content text not null,
    created_at timestamp not null,
    updated_at timestamp
);

-- +goose Down
drop table notes;
