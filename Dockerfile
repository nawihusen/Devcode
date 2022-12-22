FROM golang:1.18

RUN mkdir /app
# sesuatu yang dijalankan ketika build image/pembuatan image
WORKDIR /app

COPY ./ /app

RUN go mod tidy

RUN go build -o test

CMD ["./test"]
