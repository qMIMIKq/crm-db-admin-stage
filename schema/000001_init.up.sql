CREATE TABLE groups
(
    group_id   SERIAL UNIQUE PRIMARY KEY NOT NULL,
    group_name varchar(255)
);

INSERT INTO groups (group_name)
VALUES ('админ'),
       ('технолог'),
       ('менеджер'),
       ('оператор'),
       ('супер-админ');

CREATE TABLE plots
(
    plot_id   SERIAL UNIQUE PRIMARY KEY NOT NULL,
    plot_name VARCHAR(255)              NOT NULL
);

INSERT INTO plots (plot_name)
VALUES ('все'),
       ('фрезерный участок'),
       ('токарный участок'),
       ('фрезерный (рыжики) участок');

CREATE TABLE filters
(
    filter_id   SERIAL UNIQUE PRIMARY KEY                        NOT NULL,
    filter_name VARCHAR(255)                                     NOT NULL,
    plot_id     INT REFERENCES plots (plot_id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE users
(
    user_id   SERIAL UNIQUE PRIMARY KEY                          NOT NULL,
    user_name VARCHAR(255),
    login     VARCHAR(255),
    password  VARCHAR(255)
);

CREATE TABLE groups_description
(
    description_id serial unique primary key                          not null,
    group_id       int references groups (group_id) ON DELETE CASCADE NOT NULL,
    description    varchar(255) default 'Нет данных'
);

INSERT INTO groups_description (group_id, description)
VALUES (1, 'Может всё изменять и добавлять'),
       (2, 'Имеет доступ к части таблицы'),
       (3, 'Смотрит работу производства'),
       (4, 'Выполняет только работу по своему частку'),
       (5, 'Может всё изменять и добавлять + доступ в админку');

INSERT INTO users (user_name, login, password)
VALUES ('Никита', 'nik_admin', '');