# Как запустить проект
```
make build
```

## Создать таблицы
```
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

```

## Примеры запросов API

### USER

Регистрация пользователя:
Если под таким email уже есть пользователь в бд, то зарегистрироваться не получиться
* POST ```/auth/sign-up```
* Тело запроса
  ```json
  {
    "username": "user",
    "email": "user@gmail.com",
    "password": "123"
  }
  ```

Получение токен авторизации:
* POST ```/auth/sign-in```
* Тело запроса
  ```json
  {
    "email": "user@gmail.com",
    "password": "123"
  }
  ```
Администратором можно стать только через бд, поменяв флаг is_admin на true

### CATEGORY

Получение списка всех категорий:
Права доступа: администратор, авторизованный пользователь, гость
* GET ```/api/categories```

Получение определенной категории:
Права доступа: администратор, авторизованный пользователь, гость
* GET ```/api/categories/{id}```

Добавление новой категории:
Права доступа: администратор
Если такая категория уже есть в бд, то создать ее копию не получиться
* POST ```/api/category```
* Тело запроса
  ```json
  {
    "title": "Sport"
  }
  ```

Изменение существующей категории:
Права доступа: администратор
Если новое название категории уже существует в бд, то изменение не произойдет
* PUT ```/api/category/{id}```
* Тело запроса
  ```json
  {
    "title": "Sport"
  }
  ```

Удаление категории:
Права доступа: администратор
* DELETE ```/api/category/{id}```

### BOOK
Если на складе книги закончились, то в выборку они не попадут.
Если у книг отсутствует категория, то в выборку они не попадут.
После удаления книги, если она была у кого-то в корзине, то она пропадет из нее.

Получение списка всех книг:
Права доступа: администратор, авторизованный пользователь, гость
* GET ```/api/books```

Получение определенной книги:
Права доступа: администратор, авторизованный пользователь, гость
* GET ```/api/book/1```

Создание:
Права доступа: администратор
* POST ```/api/book```
* Тело запроса
  ```json
  {
    "title":"Dune",
    "year_publication":1965,
    "author":"Frank Herbert",
    "price":24.99,
    "qty_stock":20,
    "category_id":1
  }
  ```

Изменение:
Права доступа: администратор
* PUT ```/api/book/1```
* Тело запроса
  ```json
  {
    "title":"Dune",
    "year_publication":1965,
    "author":"Frank Herbert",
    "price":25,
    "qty_stock":20,
    "category_id":1
  }
  ```

Удаление:
Права доступа: администратор
* DELETE ```/api/book/1```

### BASKET
Книга исчезнет из корзины, если она закончилась на складе или у нее пропали категория
Если оплатить корзину, то у каждой книги из корзины будет -1 от общего кол-ва на складе

Получение списка всех книг в корзине:
Права доступа: администратор, авторизованный пользователь
* GET ```/api/basket```

Добавить книгу в корзину:
Права доступа: администратор, авторизованный пользователь
* PUT ```/api/basket/{id}```

Удалить книгу из корзины:
Права доступа: администратор, авторизованный пользователь
* DELETE ```/api/basket/{id}```

Оплатить корзину:
Права доступа: администратор, авторизованный пользователь
* POST ```/api/basket/pay```


