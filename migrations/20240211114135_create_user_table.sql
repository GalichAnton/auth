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

create table if not exists logs (
    id serial primary key,
    action text not null,
    entity_id int not null,
    created_at timestamp not null default now()
);

comment on table logs is 'Таблица с логами';
comment on column logs.id is 'Id лога';
comment on column logs.action is 'Действие над сущностью';
comment on column logs.entity_id is 'Id сущности';
comment on column logs.created_at is 'Дата создания лога';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
drop table logs;
-- +goose StatementEnd
