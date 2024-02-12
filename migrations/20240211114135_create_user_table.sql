-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION set_updated_at_from_now()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TABLE IF NOT EXISTS users (
    id serial primary key,
    name varchar(255) not null,
    email varchar(255) not null,
    password varchar(255) not null,
    role int not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);

COMMENT ON TABLE users IS 'Таблица с пользователями';
COMMENT ON COLUMN users.id IS 'Id пользователя';
COMMENT ON COLUMN users.name IS 'Имя пользователя';
COMMENT ON COLUMN users.password IS 'Пароль пользователя';
COMMENT ON COLUMN users.created_at IS 'Дата создания пользователя';
COMMENT ON COLUMN users.updated_at IS 'Дата последнего обновления пользователя';

-- ТРИГГЕР: при обновлении таблицы записывает в updated_at текущее время
DO
$$
BEGIN
        IF NOT EXISTS(SELECT 1 FROM pg_trigger WHERE tgname = 'update_users_set_update_at') THEN
CREATE TRIGGER update_user_set_update_at
    BEFORE UPDATE
    on users
    FOR EACH ROW
    EXECUTE FUNCTION
        set_updated_at_from_now();
END IF;
END
$$ LANGUAGE 'plpgsql';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
