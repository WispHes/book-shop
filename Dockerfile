# указываю базовый образ, на основе которого будет создаваться новый образ
FROM golang:1.22.2

# копирую текущий каталог своего проекта внутрь контейнера в папку /go/src/app
# если такой директории /go/src/app в контейнере не существует, то Docker автоматически создаст
# эту директория внутри контейнера
COPY . /go/src/app

# устанавливаю рабочий каталог для последующих инструкций в Dockerfile
WORKDIR /go/src/app/cmd/app

# собираю приложение из исходного кода main.go в исполняемый файл book-shop
# -o указывает на то, как исполняемый файл должен быть назван
RUN go build -o book-shop main.go

# изменяет права доступа к файлу wait-for-it.sh делая его исполняемым
RUN chmod +x wait-for-it.sh

# указывает Docker на то, что при запуске контейнера нужно запустить исполняемый файл
# CMD ["./book-shop"] - прописал в docker-compose

