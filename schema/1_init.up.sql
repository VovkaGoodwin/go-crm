BEGIN;

CREATE TABLE IF NOT EXISTS departments
(
    id          INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name        VARCHAR UNIQUE NOT NULL,
    description VARCHAR
);

CREATE TABLE IF NOT EXISTS positions
(
    id            INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name          VARCHAR UNIQUE NOT NULL,
    department_id INTEGER        NOT NULL REFERENCES departments (id) ON DELETE CASCADE,
    role          INTEGER        NOT NULL,
    description   VARCHAR
);

CREATE TABLE IF NOT EXISTS users
(
    id            INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    email         VARCHAR UNIQUE NOT NULL,
    password      VARCHAR        NOT NULL,
    first_name    VARCHAR        NOT NULL,
    position_id   INTEGER        NOT NULL REFERENCES positions (id),
    department_id INTEGER        NOT NULL REFERENCES departments (id)
);

COMMIT;