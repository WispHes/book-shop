CREATE TABLE users
(
    id          serial PRIMARY KEY NOT NULL unique,
    username    varchar(255) NOT NULL,
    email       varchar(255) NOT NULL,
    password    varchar(255) NOT NULL,
    is_admin    BOOLEAN DEFAULT false
);

create table categories
(
    id          serial PRIMARY KEY NOT NULL unique,
    title       varchar(255) NOT NULL
);

create table books
(
    id                  serial PRIMARY KEY NOT NULL unique,
    title               varchar(255) NOT NULL,
    year_publication    INT,
    author              varchar(255) NOT NULL,
    price               DECIMAL,
    qty_stock           INT,
    category_id         INT
);

create table basket
(
    id          serial PRIMARY KEY NOT NULL unique,
    userId      INT,
    book_id     INT,
    date_add    TIMESTAMP
);

INSERT INTO categories (id, title) VALUES
(1, 'Science Fiction'),
(2, 'Fantasy'),
(3, 'Mystery');


