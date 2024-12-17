BEGIN;
CREATE TABLE IF NOT EXISTS users
(
    id       INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    username VARCHAR UNIQUE NOT NULL,
    email    VARCHAR UNIQUE NOT NULL,
    password varchar        NOT NULL,
    UNIQUE (username, email)
);

INSERT INTO users (username, email, password)
VALUES ('vovka', 'admin@admin.com', '3132332d663b61272d5b5b7171388e8256ac92139282661c4475216b0b362481');

COMMIT;