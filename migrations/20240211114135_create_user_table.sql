-- +goose Up
-- +goose StatementBegin
create type role_type as enum ('unknown', 'user', 'admin');

create table if not exists roles (
    id integer primary key,
    name role_type not null
);

insert into roles (id, name)
values (0, 'unknown'), (1, 'user'), (2, 'admin');

create table if not exists role_permissions (
    id serial primary key,
    role_id integer not null references roles(id),
    permission text not null
);

insert into role_permissions (role_id, permission)
values
    ((select id from roles where name = 'user'), '/ChatV1/SendMessage'),
    ((select id from roles where name = 'admin'), '/ChatV1/Delete'),
    ((select id from roles where name = 'admin'), '/ChatV1/Create');

create table if not exists users (
    id serial primary key,
    name text not null,
    email text not null unique,
    password text not null,
    role_id integer not null references roles(id),
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
