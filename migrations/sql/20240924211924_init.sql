-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "user"
(
    id SERIAL PRIMARY KEY NOT NULL,
    login varchar NOT NULL,
    password varchar NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    CONSTRAINT user_unique UNIQUE (login)
);

CREATE TYPE secret_type AS ENUM (
    'AUTH_PAIR',
    'TEXT',
    'BINARY',
    'PAYCARD'
);

CREATE TABLE IF NOT EXISTS "secret"
(
    id SERIAL PRIMARY KEY NOT NULL,
    user_id bigint NOT NULL,
    type secret_type NOT NULL,
    name varchar NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    CONSTRAINT fk_user_secrets FOREIGN KEY(user_id) REFERENCES "user"(id)
);

CREATE TABLE IF NOT EXISTS "metadata"
(
    id SERIAL PRIMARY KEY NOT NULL,
    secret_id bigint NOT NULL,
    key varchar NOT NULL,
    value text NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    CONSTRAINT fk_secret_metadata FOREIGN KEY(secret_id) REFERENCES "secret"(id)
);

CREATE TABLE IF NOT EXISTS "login"
(
    id SERIAL PRIMARY KEY NOT NULL,
    secret_id bigint NOT NULL,
    login varchar NOT NULL,
    password varchar NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    CONSTRAINT fk_secret_metadata FOREIGN KEY(secret_id) REFERENCES "secret"(id)
);

CREATE TABLE IF NOT EXISTS "pay_card"
(
    id SERIAL PRIMARY KEY NOT NULL,
    secret_id bigint NOT NULL,
    card_number varchar NOT NULL,
    holder varchar NOT NULL,
    dates varchar NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    CONSTRAINT fk_secret_metadata FOREIGN KEY(secret_id) REFERENCES "secret"(id)
);


CREATE TABLE IF NOT EXISTS "file"
(
    id SERIAL PRIMARY KEY NOT NULL,
    secret_id bigint NOT NULL,
    path varchar NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    CONSTRAINT fk_secret_metadata FOREIGN KEY(secret_id) REFERENCES "secret"(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.login;
DROP TABLE IF EXISTS public.file;
DROP TABLE IF EXISTS public.pay_card;
DROP TABLE IF EXISTS public.metadata;
DROP TABLE IF EXISTS public.secret;
DROP TABLE IF EXISTS public.user;

DROP TYPE secret_type;
-- +goose StatementEnd