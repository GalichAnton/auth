-- +goose Up
-- +goose StatementBegin
create table if not exists users (
    id serial primary key,
    name text not null,
    email text not null,
    password text not null,
    role int not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);

comment on table users is 'Таблица с пользователями';
comment on column users.id is 'Id пользователя';
comment on column users.name is 'Имя пользователя';
comment on column users.password is 'Пароль пользователя';
comment on column users.created_at is 'Дата создания пользователя';
comment on column users.updated_at is 'Дата последнего обновления пользователя';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
