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
    user_id     INT NOT NULL,
    book_id     INT NOT NULL
);

INSERT INTO categories (id, title) VALUES
(1, 'Science Fiction'),
(2, 'Fantasy'),
(3, 'Mystery');

INSERT INTO books (id, title, year_publication, author, price, qty_stock, category_id) VALUES
(1, 'Dune', 1965, 'Frank Herbert', 24.99, 50, 1),
(2, 'Harry Potter and the Philosophers Stone', 1997, 'J.K. Rowling', 19.99, 1, 2),
(3, 'The Da Vinci Code', 2003, 'Dan Brown', 29.99, 1, 3);

INSERT INTO basket (user_id, book_id) VALUES
(1, 1),
(2, 2),
(1, 2);







INSERT INTO users (id, username, email, password, is_admin) VALUES
(1, 'admin', 'admin@gmail.com', 'admin123', true),
(2, 'user1', 'user1@gmail.com', 'user123', false),
(3, 'user2', 'user2@gmail.com', 'user123', false);


