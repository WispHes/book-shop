CREATE TABLE users
(
    id          serial PRIMARY KEY NOT NULL unique,
    username    varchar(255) NOT NULL unique,
    email       varchar(255) NOT NULL unique,
    password    varchar(255) NOT NULL,
    is_admin    BOOLEAN DEFAULT false
);

create table categories
(
    id          serial PRIMARY KEY NOT NULL unique,
    title       varchar(255) NOT NULL unique
);

create table books
(
    id                  serial PRIMARY KEY NOT NULL unique,
    title               varchar(255) NOT NULL unique,
    year_publication    INT,
    author              varchar(255) NOT NULL,
    price               DECIMAL,
    qty_stock           INT,
    category_id         INT
);

create table basket
(
    user_id     INT NOT NULL,
    book_id     INT NOT NULL unique
);


